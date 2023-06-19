package choose

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	navigatorTUI "github.com/anondigriz/mogan-mini/internal/tui/group/navigator"
	chooseTUI "github.com/anondigriz/mogan-mini/internal/tui/table/choose"
)

type AllowArgs struct {
	Group     bool
	Parameter bool
	Rule      bool
}

type ShowArgs struct {
	Parameter bool
	Rule      bool
}

type Choose struct {
	lg         *zap.Logger
	Parameters map[string]kbEnt.Parameter
	Rules      map[string]kbEnt.Rule
	Root       *navigatorTUI.Group
}

func New(lg *zap.Logger) *Choose {
	c := &Choose{
		lg: lg,
	}
	return c
}

func (c Choose) ChooseTUI(allow AllowArgs, show ShowArgs) (string, error) {
	mt := chooseTUI.New(
		navigatorTUI.New(
			c.Root,
			navigatorTUI.AllowArgs{
				Group:     allow.Group,
				Parameter: allow.Parameter,
				Rule:      allow.Rule,
			},
			navigatorTUI.ShowArgs{
				Parameter: show.Parameter,
				Rule:      show.Rule,
			},
		))

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

	nav, ok := result.Nav.(*navigatorTUI.Navigator)
	if !ok {
		err := fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Error(err.Error())
		return "", err
	}

	return nav.ChosenUUID, nil
}
