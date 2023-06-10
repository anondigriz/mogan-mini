package description

// A simple program demonstrating the textarea component from the Bubbles
// component library.

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

type Model struct {
	keys      keyMap
	textarea  textarea.Model
	help      help.Model
	Info      kbEnt.BaseInfo
	IsEdited  bool
	IsQuitted bool
	Err       error
}

func New(info kbEnt.BaseInfo) Model {
	m := Model{
		keys: keys,
		help: help.New(),
		Info: info,
	}
	m.buildForm()
	return m
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			m.IsQuitted = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Save):
			m.IsEdited = true
			m.Info.Description = m.textarea.Value()
			m.Info.ModifiedDate = time.Now().UTC()
			return m, tea.Quit
		case key.Matches(msg, m.keys.FocusMode):
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.IsQuitted || m.IsEdited {
		return ""
	}

	helpView := m.help.View(m.keys)
	textareaView := m.textarea.View()

	height := 2
	if strings.Count(helpView, "\n") > 0 {
		height = 1
	}

	return textareaView + strings.Repeat("\n", height) + helpView
}
