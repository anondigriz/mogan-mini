package textinput

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type Model struct {
	Question  string
	TextInput textinput.Model
	IsQuitted bool
	Err       error
}

func New(question, placeholder string) Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20

	return Model{
		Question:  question,
		TextInput: ti,
		Err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.IsQuitted = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.Err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.Question,
		m.TextInput.View(),
		"(esc or ctrl+c to quit)",
	) + "\n"
}
