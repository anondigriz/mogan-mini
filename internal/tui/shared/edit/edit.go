package edit

import (
	entKB "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type DescriptionModel struct {
	Description string
	IsEdited    bool
}

type Model struct {
	BaseInfo    BaseInfoModel
	Description DescriptionModel
	IsQuitting  bool
}

func New(bi entKB.BaseInfo) Model {
	m := Model{
		BaseInfo: newBaseInfo(bi),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.baseInfoInit()
}

func (m Model) baseInfoInit() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.IsQuitting = true
			return m, tea.Quit
		}
	}
	if !m.BaseInfo.IsEdited {
		//TODO
	}

	return m.updateBaseInfo(msg)
}

func (m Model) View() string {
	if !m.BaseInfo.IsEdited {
		//TODO
	}

	return m.BaseInfo.view()
}
