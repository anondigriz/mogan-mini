package choose

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	editDesTUI "github.com/anondigriz/mogan-mini/internal/tui/baseinfo/edit/description"
	editISTUI "github.com/anondigriz/mogan-mini/internal/tui/baseinfo/edit/idshortname"
)

type Edit struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *Edit {
	c := &Edit{
		lg: lg,
	}
	return c
}

func (e Edit) EditTUI(info kbEnt.BaseInfo) (kbEnt.BaseInfo, error) {
	messages.PrintEditIDShortName()
	info, err := e.editIdShortNameTUI(info)
	if err != nil {
		e.lg.Error(errMsgs.EditTUIIDShortNameFail, zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}

	messages.PrintEditDescription()
	info, err = e.editDescriptionTUI(info)
	if err != nil {
		e.lg.Error(errMsgs.EditTUIDescriptionFail, zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}

	return info, nil
}

func (e Edit) editIdShortNameTUI(info kbEnt.BaseInfo) (kbEnt.BaseInfo, error) {
	mt := editISTUI.New(info)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		e.lg.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}
	result, ok := m.(editISTUI.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		e.lg.Error(err.Error(), zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}

	if result.IsQuitted || !result.IsEdited {
		err = fmt.Errorf(errMsgs.BaseInfoNotEdited)
		e.lg.Error(err.Error(), zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}
	return result.Info, nil
}

func (e Edit) editDescriptionTUI(info kbEnt.BaseInfo) (kbEnt.BaseInfo, error) {
	mt := editDesTUI.New(info)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		e.lg.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}
	result, ok := m.(editDesTUI.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		e.lg.Error(err.Error(), zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}

	if result.IsQuitted || !result.IsEdited {
		err = fmt.Errorf(errMsgs.BaseInfoNotEdited)
		e.lg.Error(err.Error(), zap.Error(err))
		return kbEnt.BaseInfo{}, err
	}
	return result.Info, nil
}
