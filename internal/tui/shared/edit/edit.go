package edit

import (
	entKB "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	BaseInfo    baseInfoModel
	Description descriptionModel
	IsQuitted   bool
}

func New(bi entKB.BaseInfo, description string) Model {
	m := Model{
		BaseInfo:    newBaseInfoModel(bi),
		Description: newDescriptionModel(description),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.BaseInfo.init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "ctrl+c" || k == "esc" {
			m.IsQuitted = true
			return m, tea.Quit
		}
	}
	if m.BaseInfo.IsEdited {
		return m.updateDescription(msg)
	}

	return m.updateBaseInfo(msg)
}

func (m Model) View() string {
	if m.BaseInfo.IsEdited {
		return m.Description.view()
	}

	return m.BaseInfo.view()
}
