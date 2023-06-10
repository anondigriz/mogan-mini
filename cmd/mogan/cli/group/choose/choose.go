package choose

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	chooseTUI "github.com/anondigriz/mogan-mini/internal/tui/group/choose"
)

type Choose struct {
	lg         *zap.Logger
	Parameters map[string]kbEnt.Parameter
	Rules      map[string]kbEnt.Rule
	Root       *chooseTUI.Group
}

func New(lg *zap.Logger) *Choose {
	c := &Choose{
		lg: lg,
	}
	return c
}

func (c Choose) ChooseTUI(allow chooseTUI.AllowArgs, show chooseTUI.ShowArgs) (string, error) {
	mt := chooseTUI.New(c.Root, allow, show)

	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return "", err
	}
	result, ok := m.(chooseTUI.Model)
	if !ok {
		err := fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Error(err.Error())
		return "", err
	}

	if result.IsQuitted {
		err := fmt.Errorf(errMsgs.ObjectWasNotChosen)
		c.lg.Error(err.Error())
		return "", err
	}

	return result.ChosenUUID, nil
}
