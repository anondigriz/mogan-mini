package description

import (
	"github.com/charmbracelet/bubbles/textarea"
)

func (m *Model) buildForm() {
	ta := textarea.New()
	ta.Placeholder = "This is a very important object..."
	ta.Focus()
	ta.SetValue(m.Info.Description)
	m.textarea = ta
}
