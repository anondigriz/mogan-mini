package listshow

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	List []string
}

func New(list []string) Model {
	return Model{
		List: list,
	}
}

func (lm Model) Init() tea.Cmd {
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

func (lm Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return lm, tea.Quit
	}
	return lm, nil
}

func (lm Model) View() string {
	return ""
}
