package edit

// A simple program demonstrating the textarea component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type descriptionModel struct {
	textarea    textarea.Model
	err         error
	Description string
	IsEdited    bool
}

type errMsg error

func newDescriptionModel(description string) descriptionModel {
	ti := textarea.New()
	ti.Placeholder = "This is a very important object..."
	ti.Focus()

	return descriptionModel{
		Description: description,
		textarea:    ti,
		err:         nil,
	}
}

func (m descriptionModel) init() tea.Cmd {
	return textarea.Blink
}

func (m Model) updateDescription(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.Description.textarea.Focused() {
				m.Description.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.Description.textarea.Focused() {
				cmd = m.Description.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.Description.err = msg
		return m, nil
	}

	m.Description.textarea, cmd = m.Description.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m descriptionModel) view() string {
	return fmt.Sprintf(
		"Edit the description.\n\n%s\n\n%s",
		m.textarea.View(),
		"(esc or ctrl+c to quit)",
	) + "\n\n"
}
