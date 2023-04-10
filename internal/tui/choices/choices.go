package choices

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	cursor    int
	question  string
	choices   []string
	Choice    string
	IsQuitted bool
}

func New(question string, choices []string) Model {
	m := Model{
		question: question,
		choices:  choices,
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.IsQuitted = true
			return m, tea.Quit
		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = m.choices[m.cursor]
			return m, tea.Quit
		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.question)
	s.WriteString("\n\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(esc or ctrl+c to quit)\n")

	return s.String()
}
