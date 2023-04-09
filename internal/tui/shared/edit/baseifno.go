package edit

import (
	"fmt"
	"strings"
	"time"

	entKB "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Copy().Render("[ Next ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Next"))
)

type BaseInfoModel struct {
	focusIndex   int
	inputs       []textinput.Model
	ID           string
	ShortName    string
	ModifiedDate time.Time
	IsEdited     bool
}

func newBaseInfo(bi entKB.BaseInfo) BaseInfoModel {
	m := BaseInfoModel{
		inputs:       make([]textinput.Model, 2),
		ID:           bi.ID,
		ShortName:    bi.ShortName,
		ModifiedDate: bi.ModifiedDate,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle

		switch i {
		case 0:
			t.Placeholder = "ID"
			t.CharLimit = 40
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.SetValue(bi.ID)
		case 1:
			t.Placeholder = "Short name"
			t.CharLimit = 50
			t.SetValue(bi.ShortName)
		}

		m.inputs[i] = t
	}

	return m
}

func (m Model) updateBaseInfo(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.BaseInfo.focusIndex == len(m.BaseInfo.inputs) {
				m.BaseInfo.IsEdited = true
				m.BaseInfo.ID = m.BaseInfo.inputs[0].Value()
				m.BaseInfo.ShortName = m.BaseInfo.inputs[1].Value()
				m.BaseInfo.ModifiedDate = time.Now().UTC()

				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.BaseInfo.focusIndex--
			} else {
				m.BaseInfo.focusIndex++
			}

			if m.BaseInfo.focusIndex > len(m.BaseInfo.inputs) {
				m.BaseInfo.focusIndex = 0
			} else if m.BaseInfo.focusIndex < 0 {
				m.BaseInfo.focusIndex = len(m.BaseInfo.inputs)
			}

			cmds := make([]tea.Cmd, len(m.BaseInfo.inputs))
			for i := 0; i <= len(m.BaseInfo.inputs)-1; i++ {
				if i == m.BaseInfo.focusIndex {
					// Set focused state
					cmds[i] = m.BaseInfo.inputs[i].Focus()
					m.BaseInfo.inputs[i].PromptStyle = focusedStyle
					m.BaseInfo.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.BaseInfo.inputs[i].Blur()
				m.BaseInfo.inputs[i].PromptStyle = noStyle
				m.BaseInfo.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.BaseInfo.updateInputs(msg)

	return m, cmd
}

func (m *BaseInfoModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m BaseInfoModel) view() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n%s\n", *button, "(esc or q to quit)")

	return b.String()
}
