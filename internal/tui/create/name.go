package create

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type NameModel struct {
	TextInput textinput.Model
	Err       error
}

func InitialNameModel() NameModel {
	ti := textinput.New()
	ti.Placeholder = "Awesome knowledge base"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20

	return NameModel{
		TextInput: ti,
		Err:       nil,
	}
}

func (n NameModel) Init() tea.Cmd {
	return textinput.Blink
}

func (n NameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (n NameModel) View() string {
	return fmt.Sprintf(
		"What is the name of the knowledge base?\n\n%s\n\n%s",
		n.TextInput.View(),
		"(esc to quit)",
	) + "\n"
}
