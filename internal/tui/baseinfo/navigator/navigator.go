package navigator

import (
	"strconv"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type Navigator struct {
	info          []kbEnt.BaseInfo
	TimeFormat    string
	TableHeight   int
	BaseStyle     lipgloss.Style
	HeaderStyle   lipgloss.Style
	SelectedStyle lipgloss.Style
	Columns       []table.Column
	ChosenUUID    string
}

func New(info []kbEnt.BaseInfo) *Navigator {
	nav := &Navigator{
		info:        info,
		TimeFormat:  "02.01.2006 15:04:05",
		TableHeight: 7,
		BaseStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")),
		HeaderStyle: table.DefaultStyles().Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false),
		SelectedStyle: table.DefaultStyles().Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false),
		Columns: []table.Column{
			{Title: "#", Width: 4},
			{Title: "Short name", Width: 15},
			{Title: "UUID", Width: 10},
			{Title: "ID", Width: 10},
			{Title: "Modified", Width: 20},
		},
	}
	return nav
}

func (nav *Navigator) Build() table.Model {
	rows := make([]table.Row, 0, len(nav.info))
	for i, v := range nav.info {
		rows = append(rows, table.Row{
			strconv.Itoa(i + 1),
			v.ShortName,
			v.UUID,
			v.ID,
			v.ModifiedDate.Local().Format(nav.TimeFormat)})
	}

	t := table.New(
		table.WithColumns(nav.Columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(nav.TableHeight),
	)
	s := table.DefaultStyles()
	s.Header = nav.HeaderStyle
	s.Selected = nav.SelectedStyle
	t.SetStyles(s)
	return t
}

func (nav *Navigator) Render(t table.Model) string {
	nav.BaseStyle.Render(t.View())
	return nav.BaseStyle.Render(t.View())
}

func (nav *Navigator) Open(r table.Row) bool {
	return false
}

func (nav *Navigator) Back() bool {
	return false
}

func (nav *Navigator) Choose(r table.Row) bool {
	nav.ChosenUUID = r[2]
	return true
}
