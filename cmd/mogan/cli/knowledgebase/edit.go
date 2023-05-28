package knowledgebase

import (
	"context"
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	editTui "github.com/anondigriz/mogan-mini/internal/tui/shared/edit"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
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
		err := fmt.Errorf(errMsgs.KnowledgeBaseNotChosen)
		e.lg.Zap.Error(err.Error())
		messages.PrintFail(errMsgs.KnowledgeBaseNotChosen)
		return err
	}

	st := kbsSt.New(e.lg.Zap, e.cfg.WorkspaceDir)
	kbsu := kbsUC.New(e.lg.Zap, st)
	kb, err := kbsu.GetKnowledgeBase(e.cfg.CurrentKnowledgeBase.UUID)
	if err != nil {
		e.lg.Zap.Error(errMsgs.GetKnowledgeBase, zap.Error(err))
		messages.PrintFail(errMsgs.GetKnowledgeBase)
		return err
	}

	updated, err := e.editTUIKnowledgeBase(cmd.Context(), kb)
	if err != nil {
		e.lg.Zap.Error(errMsgs.EditTUIKnowledgeBase, zap.Error(err))
		messages.PrintFail(errMsgs.EditTUIKnowledgeBase)
		return err
	}

	return e.commitChanges(cmd.Context(), kbsu, updated)
}

func (e Edit) editTUIKnowledgeBase(ctx context.Context, previous kbEnt.KnowledgeBase) (kbEnt.KnowledgeBase, error) {
	mt := editTui.New(previous.BaseInfo, previous.Description)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		e.lg.Zap.Error(errMsgs.RunTUIProgram, zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	result, ok := m.(editTui.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}

	if result.IsQuitted || !result.BaseInfo.IsEdited || !result.Description.IsEdited {
		err = fmt.Errorf(errMsgs.KnowledgeBaseWasNotEdited)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}

	if result.BaseInfo.ID == "" {
		err = fmt.Errorf(errMsgs.IDIsEmpty)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}
	if result.BaseInfo.ShortName == "" {
		err = fmt.Errorf(errMsgs.ShortNameIsEmpty)
		e.lg.Zap.Error(err.Error())
		return kbEnt.KnowledgeBase{}, err
	}

	var updated kbEnt.KnowledgeBase = previous
	updated.BaseInfo.ID = result.BaseInfo.ID
	updated.BaseInfo.ShortName = result.BaseInfo.ShortName
	updated.BaseInfo.ModifiedDate = result.BaseInfo.ModifiedDate
	updated.Description = result.Description.Description

	return updated, nil
}

func (e Edit) commitChanges(ctx context.Context, kbsu *kbsUC.KnowledgeBases, updated kbEnt.KnowledgeBase) error {
	messages.PrintReceivedNewEntityInfo()

	err := kbsu.UpdateKnowledgeBase(updated)
	if err != nil {
		e.lg.Zap.Error(errMsgs.UpdateKnowledgeBase, zap.Error(err))
		messages.PrintFail(errMsgs.UpdateKnowledgeBase)
		return err
	}
	messages.PrintChangesAccepted()
	return nil
}
