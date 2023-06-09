package choose

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	chooseTui "github.com/anondigriz/mogan-mini/internal/tui/group/choose"
)

type Choose struct {
	lg         *zap.Logger
	Parameters map[string]kbEnt.Parameter
	Rules      map[string]kbEnt.Rule
	Root       *chooseTui.Group
}

func New(lg *zap.Logger) *Choose {
	c := &Choose{
		lg: lg,
	}
	return c
}

func (c Choose) ChooseGroup() (string, error) {
	messages.PrintChooseGroup()
	uuid, err := c.ChooseTUI(
		chooseTui.AllowArgs{Group: true, Parameter: false, Rule: false},
		chooseTui.ShowArgs{Parameter: true, Rule: true})
	if err != nil {
		c.lg.Error(errMsgs.ChooseTUIGroupFail, zap.Error(err))
		return "", err
	}
	return uuid, nil
}

func (c Choose) ChooseTUI(allow chooseTui.AllowArgs, show chooseTui.ShowArgs) (string, error) {
	mt := chooseTui.New(c.Root, allow, show)

	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return "", err
	}
	result, ok := m.(chooseTui.Model)
	if !ok {
		err := fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Error(err.Error())
		return "", err
	}

	if result.IsQuitted {
		err := fmt.Errorf(errMsgs.KnowledgeBaseWasNotChosen)
		c.lg.Error(err.Error())
		return "", err
	}

	return result.ChosenUUID, nil
}
