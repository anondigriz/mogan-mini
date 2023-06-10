package idshortname

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func (m *Model) buildForm() {
	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle

		switch i {
		case 0:
			t.Placeholder = "ID"
			t.CharLimit = 40
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.SetValue(m.Info.ID)
		case 1:
			t.Placeholder = "Short name"
			t.CharLimit = 50
			t.SetValue(m.Info.ShortName)
		}

		m.inputs[i] = t
	}
}
