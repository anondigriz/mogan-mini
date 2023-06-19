package choose

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Navigator interface {
	Build() table.Model
	Render(t table.Model) string
	Open(r table.Row) bool
	Back() bool
	Choose(r table.Row) bool
}

type Model struct {
	keys      keyMap
	table     table.Model
	help      help.Model
	Nav       Navigator
	IsQuitted bool
	IsChosen  bool
}

func New(nav Navigator) Model {
	m := Model{
		keys: keys,
		help: help.New(),
		Nav:  nav,
	}
	m.table = m.Nav.Build()
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
			if m.Nav.Back() {
				m.table = m.Nav.Build()
			}
		case key.Matches(msg, m.keys.Quit):
			m.IsQuitted = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Choose):
			if m.Nav.Choose(m.table.SelectedRow()) {
				m.IsChosen = true
				return m, tea.Quit
			}
		case key.Matches(msg, m.keys.Open):
			if m.Nav.Open(m.table.SelectedRow()) {
				m.table = m.Nav.Build()
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
	tableView := m.Nav.Render(m.table)

	height := 2
	if strings.Count(helpView, "\n") > 0 {
		height = 1
	}
	return tableView + strings.Repeat("\n", height) + helpView
}
