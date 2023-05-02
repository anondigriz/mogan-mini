package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/logger"
	editTui "github.com/anondigriz/mogan-mini/internal/tui/shared/edit"
	kbUseCase "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase"
)

type Edit struct {
	lg  *logger.Logger
	cfg *config.Config
	Cmd *cobra.Command
}

func NewEdit(lg *logger.Logger, cfg *config.Config) *Edit {
	e := &Edit{
		lg:  lg,
		cfg: cfg,
	}

	e.Cmd = &cobra.Command{
		Use:   "edit",
		Short: "Editing the database",
		Long:  `Editing basic information about the knowledge base`,
		RunE:  e.runE,
	}
	return e
}

func (e *Edit) Init() {
	cobra.OnInitialize(e.initConfig)
}

func (e *Edit) initConfig() {
}

func (e *Edit) runE(cmd *cobra.Command, args []string) error {
	if e.cfg.CurrentKnowledgeBase.UUID == "" {
		err := fmt.Errorf(errors.KnowledgeBaseNotChosenErrMsg)
		e.lg.Zap.Error(err.Error())
		messages.PrintFail(errors.KnowledgeBaseNotChosenErrMsg)
		return err
	}

	kbu := kbUseCase.New(e.lg.Zap, *e.cfg)
	kb, err := kbu.Get(cmd.Context(), e.cfg.CurrentKnowledgeBase.UUID)
	if err != nil {
		e.lg.Zap.Error(errors.GetKnowledgeBaseErrMsg, zap.Error(err))
		messages.PrintFail(errors.GetKnowledgeBaseErrMsg)
		return err
	}

	updated, err := e.editTUIKnowledgeBase(cmd.Context(), kb)
	if err != nil {
		e.lg.Zap.Error(errors.EditTUIKnowledgeBaseErrMsg, zap.Error(err))
		messages.PrintFail(errors.EditTUIKnowledgeBaseErrMsg)
		return err
	}

	return e.commitChanges(cmd.Context(), kbu, updated)
}

func (e Edit) editTUIKnowledgeBase(ctx context.Context, previous kbEnt.KnowledgeBase) (kbEnt.KnowledgeBase, error) {
	mt := editTui.New(previous.BaseInfo, previous.ExtraData.Description)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		e.lg.Zap.Error(errors.RunTUIProgramErrMsg, zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	result, ok := m.(editTui.Model)
	if !ok {
		err = fmt.Errorf(errors.ReceivedResponseWasNotExpectedErrMsg)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}

	if result.IsQuitted || !result.BaseInfo.IsEdited || !result.Description.IsEdited {
		err = fmt.Errorf(errors.KnowledgeBaseWasNotEditedErrMsg)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}

	if result.BaseInfo.ID == "" {
		err = fmt.Errorf(errors.IDIsEmptyErrMsg)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}
	if result.BaseInfo.ShortName == "" {
		err = fmt.Errorf(errors.ShortNameIsEmptyErrMsg)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}

	var updated kbEnt.KnowledgeBase = previous
	updated.BaseInfo.ID = result.BaseInfo.ID
	updated.BaseInfo.ShortName = result.BaseInfo.ShortName
	updated.BaseInfo.ModifiedDate = result.BaseInfo.ModifiedDate
	updated.ExtraData.Description = result.Description.Description

	return updated, nil
}

func (e Edit) commitChanges(ctx context.Context, kbu *kbUseCase.KnowledgeBase, updated kbEnt.KnowledgeBase) error {
	messages.PrintRecivedNewEntityInfo()

	err := kbu.Update(ctx, updated)
	if err != nil {
		e.lg.Zap.Error(errors.UpdateKnowledgeBaseErrMsg, zap.Error(err))
		messages.PrintFail(errors.UpdateKnowledgeBaseErrMsg)
		return err
	}
	messages.PrintChangesAccepted()
	return nil
}
