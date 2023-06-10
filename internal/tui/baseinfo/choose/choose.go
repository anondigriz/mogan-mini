package choose

import (
	"strings"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	timeFormat  = "02.01.2006 15:04:05"
	tableHeight = 7
	baseStyle   = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
	columns = []table.Column{
		{Title: "#", Width: 4},
		{Title: "Short name", Width: 15},
		{Title: "UUID", Width: 10},
		{Title: "ID", Width: 10},
		{Title: "Modified", Width: 20},
	}
)

type Model struct {
	keys       keyMap
	table      table.Model
	help       help.Model
	info       []kbEnt.BaseInfo
	IsQuitted  bool
	IsChosen   bool
	ChosenUUID string
}

func New(info []kbEnt.BaseInfo) Model {
	m := Model{
		keys: keys,
		help: help.New(),
		info: info,
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
		case key.Matches(msg, m.keys.Quit):
			m.IsQuitted = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Choose):
			m.ChosenUUID = m.table.SelectedRow()[1]
			m.IsChosen = true
			return m, tea.Quit
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
