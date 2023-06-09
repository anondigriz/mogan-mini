package choose

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	timeFormat = "02.01.2006 15:04:05"
)

var (
	tableHeight = 7
	baseStyle   = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
	columns = []table.Column{
		{Title: "Short name", Width: 15},
		{Title: "UUID", Width: 10},
		{Title: "ID", Width: 10},
		{Title: "Type", Width: 4},
		{Title: "Modified", Width: 20},
	}
)

type AllowArgs struct {
	Group     bool
	Parameter bool
	Rule      bool
}

type ShowArgs struct {
	Parameter bool
	Rule      bool
}

type Model struct {
	keys             keyMap
	table            table.Model
	help             help.Model
	root             *Group
	IsQuitted        bool
	IsChosen         bool
	ChosenUUID       string
	ChosenObjectType TypeObject
	Allow            AllowArgs
	Show             ShowArgs
}

func New(root *Group, allow AllowArgs, show ShowArgs) Model {
	m := Model{
		keys:  keys,
		help:  help.New(),
		root:  root,
		Allow: allow,
		Show:  show,
	}
	m.buildTable()

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Back):
			if m.root.Parent != nil {
				m.root = m.root.Parent
				m.buildTable()
				break
			}
		case key.Matches(msg, m.keys.Quit):
			m.IsQuitted = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Choose):
			t := m.table.SelectedRow()[3]
			uuid := m.table.SelectedRow()[1]
			if t == groupTypeShortName && m.Allow.Group {
				m.ChosenUUID = uuid
				m.ChosenObjectType = GroupType
				m.IsChosen = true
				return m, tea.Quit
			}
			if t == parameterTypeShortName && m.Allow.Parameter {
				m.ChosenUUID = uuid
				m.ChosenObjectType = ParameterType
				m.IsChosen = true
				return m, tea.Quit
			}
			if t == ruleTypeShortName && m.Allow.Rule {
				m.ChosenUUID = uuid
				m.ChosenObjectType = RuleType
				m.IsChosen = true
				return m, tea.Quit
			}
		case key.Matches(msg, m.keys.Open):
			shortName := m.table.SelectedRow()[0]
			t := m.table.SelectedRow()[3]
			uuid := m.table.SelectedRow()[1]

			if shortName == rootShortName && m.root.Parent != nil {
				m.root = m.root.Parent
				m.buildTable()
				break
			}
			if t != groupTypeShortName {
				break
			}

			for _, v := range m.root.Groups {
				if v.UUID != uuid {
					continue
				}
				m.root = v
				m.buildTable()
				break
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.IsQuitted || m.IsChosen {
		return ""
	}

	helpView := m.help.View(m.keys)
	tableView := baseStyle.Render(m.table.View())
	height := tableHeight + 5 - strings.Count(tableView, "\n") - strings.Count(helpView, "\n")

	return tableView + strings.Repeat("\n", height) + helpView
}
