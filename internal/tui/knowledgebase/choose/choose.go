package choose

import (
	"strconv"

	"github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	table  table.Model
	Choice string
}

func New(kbs []knowledgebase.KnowledgeBase) Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "UUID", Width: 10},
		{Title: "Short name", Width: 15},
		{Title: "Modified", Width: 20},
	}

	rows := make([]table.Row, 0, len(kbs))

	for i, v := range kbs {
		rows = append(rows, table.Row{strconv.Itoa(i + 1), v.UUID, v.ShortName, v.ModifiedDate.Format("02.01.2006 15:04:05")})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := Model{
		table: t,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.Choice = m.table.SelectedRow()[1]
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.table.View()) + "\n(q to quit)\n"
}
