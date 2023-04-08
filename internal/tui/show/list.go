package show

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	List []string
}

func (lm ListModel) Init() tea.Cmd {
	var batch []tea.Cmd
	for _, v := range lm.List {
		batch = append(batch, tea.Println(v))
	}

	return tea.Sequence(
		tea.Batch(
			batch...,
		),
		tea.Quit,
	)
}

func (lm ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return lm, tea.Quit
	}
	return lm, nil
}

func (lm ListModel) View() string {
	return ""
}
