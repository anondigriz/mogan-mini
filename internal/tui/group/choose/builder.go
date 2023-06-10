package choose

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) buildTable() {
	rowsCount := len(m.root.Groups) + len(m.root.Parameters) + len(m.root.Rules)
	if m.root.Parent != nil {
		rowsCount += 1
	}
	rows := make([]table.Row, 0, rowsCount)

	if m.root.Parent != nil {
		rows = append(rows, table.Row{
			rootShortName,
			"",
			"",
			"",
			""})
	}

	for _, v := range m.root.Groups {
		rows = append(rows, table.Row{
			fmt.Sprintf("/%s", v.ShortName),
			v.UUID,
			v.ID,
			groupTypeShortName,
			v.ModifiedDate.Local().Format(timeFormat)})
	}

	if m.Show.Parameter {
		for _, v := range m.root.Parameters {
			rows = append(rows, table.Row{
				v.ShortName,
				v.UUID,
				v.ID,
				parameterTypeShortName,
				v.ModifiedDate.Local().Format(timeFormat)})
		}
	}

	if m.Show.Rule {
		for _, v := range m.root.Rules {
			rows = append(rows, table.Row{
				v.ShortName,
				v.UUID,
				v.ID,
				ruleTypeShortName,
				v.ModifiedDate.Local().Format(timeFormat)})
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(tableHeight),
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
	m.table = t
}
