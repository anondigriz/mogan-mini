package navigator

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
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

type Navigator struct {
	root             *Group
	Allow            AllowArgs
	Show             ShowArgs
	TimeFormat       string
	TableHeight      int
	BaseStyle        lipgloss.Style
	HeaderStyle      lipgloss.Style
	SelectedStyle    lipgloss.Style
	Columns          []table.Column
	ChosenUUID       string
	ChosenObjectType TypeObject
}

func New(root *Group, allow AllowArgs, show ShowArgs) *Navigator {
	nav := &Navigator{
		root:        root,
		Allow:       allow,
		Show:        show,
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
			{Title: "Short name", Width: 15},
			{Title: "UUID", Width: 10},
			{Title: "ID", Width: 10},
			{Title: "Type", Width: 4},
			{Title: "Modified", Width: 20},
		},
	}
	return nav
}

func (nav Navigator) Build() table.Model {
	rowsCount := len(nav.root.Groups) + len(nav.root.Parameters) + len(nav.root.Rules)
	if nav.root.Parent != nil {
		rowsCount += 1
	}
	rows := make([]table.Row, 0, rowsCount)

	rows = nav.AddParentGroup(rows)
	rows = nav.AddGroups(rows)
	rows = nav.AddParameters(rows)
	rows = nav.AddRules(rows)

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

func (nav Navigator) AddParentGroup(rows []table.Row) []table.Row {
	if nav.root.Parent == nil {
		return rows
	}

	rows = append(rows, table.Row{
		rootShortName,
		"",
		"",
		"",
		""})
	return rows
}

func (nav Navigator) AddGroups(rows []table.Row) []table.Row {
	for _, v := range nav.root.Groups {
		rows = append(rows, table.Row{
			fmt.Sprintf("/%s", v.ShortName),
			v.UUID,
			v.ID,
			groupTypeShortName,
			v.ModifiedDate.Local().Format(nav.TimeFormat)})
	}
	return rows
}

func (nav Navigator) AddParameters(rows []table.Row) []table.Row {
	if !nav.Show.Parameter {
		return rows
	}
	for _, v := range nav.root.Parameters {
		rows = append(rows, table.Row{
			v.ShortName,
			v.UUID,
			v.ID,
			parameterTypeShortName,
			v.ModifiedDate.Local().Format(nav.TimeFormat)})
	}
	return rows
}

func (nav Navigator) AddRules(rows []table.Row) []table.Row {
	if !nav.Show.Rule {
		return rows
	}
	for _, v := range nav.root.Rules {
		rows = append(rows, table.Row{
			v.ShortName,
			v.UUID,
			v.ID,
			ruleTypeShortName,
			v.ModifiedDate.Local().Format(nav.TimeFormat)})
	}
	return rows
}

func (nav Navigator) Render(t table.Model) string {
	nav.BaseStyle.Render(t.View())
	return nav.BaseStyle.Render(t.View())
}

func (nav *Navigator) Open(r table.Row) bool {
	shortName := r[0]
	t := r[3]
	uuid := r[1]

	if shortName == rootShortName && nav.root.Parent != nil {
		nav.root = nav.root.Parent
		return true
	}
	if t != groupTypeShortName {
		return false
	}

	for _, v := range nav.root.Groups {
		if v.UUID != uuid {
			continue
		}
		nav.root = v
		return true
	}
	return false
}

func (nav *Navigator) Back() bool {
	if nav.root.Parent == nil {
		return false
	}
	nav.root = nav.root.Parent
	return true
}

func (nav *Navigator) Choose(r table.Row) bool {
	t := r[3]
	nav.ChosenUUID = r[1]
	switch {
	case t == groupTypeShortName && nav.Allow.Group:
		nav.ChosenObjectType = GroupType
		return true
	case t == parameterTypeShortName && nav.Allow.Parameter:
		nav.ChosenObjectType = ParameterType
		return true
	case t == ruleTypeShortName && nav.Allow.Rule:
		nav.ChosenObjectType = RuleType
		return true
	}

	nav.ChosenUUID = ""
	return false
}
