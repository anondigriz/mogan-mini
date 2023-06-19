package choose

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	navigatorTUI "github.com/anondigriz/mogan-mini/internal/tui/baseinfo/navigator"
	chooseTUI "github.com/anondigriz/mogan-mini/internal/tui/table/choose"
)

type Choose struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *Choose {
	c := &Choose{
		lg: lg,
	}
	return c
}

func (c Choose) ChooseTUI(info []kbEnt.BaseInfo) (string, error) {
	mt := chooseTUI.New(navigatorTUI.New(info))
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return "", err
	}
	result, ok := m.(chooseTUI.Model)
	if !ok {
		err := fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Error(err.Error(), zap.Error(err))
		return "", err
	}

	if result.IsQuitted {
		err := fmt.Errorf(errMsgs.ObjectWasNotChosen)
		c.lg.Error(err.Error(), zap.Error(err))
		return "", err
	}

	nav, ok := result.Nav.(*navigatorTUI.Navigator)
	if !ok {
		err := fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Error(err.Error())
		return "", err
	}

	return nav.ChosenUUID, nil
}
