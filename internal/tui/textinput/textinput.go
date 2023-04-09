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

func (n Model) Init() tea.Cmd {
	return textinput.Blink
}

func (n Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return n, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		n.Err = msg
		return n, nil
	}

	n.TextInput, cmd = n.TextInput.Update(msg)
	return n, cmd
}

func (n Model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		n.Question,
		n.TextInput.View(),
		"(esc to quit)",
	) + "\n"
}
