package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/logger"
	editTui "github.com/anondigriz/mogan-mini/internal/tui/shared/edit"
	kbManagement "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management"
)

type Edit struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewEdit(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Edit {
	e := &Edit{
		lg:  lg,
		vp:  vp,
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
		e.lg.Zap.Error(err.Error(), zap.Error(err))
		return err
	}

	man := kbManagement.New(e.lg.Zap, *e.cfg)
	kb, err := man.Get(cmd.Context(), e.cfg.CurrentKnowledgeBase.UUID)
	if err != nil {
		e.lg.Zap.Error(errors.GetKnowledgeBaseErrMsg, zap.Error(err))
		fmt.Printf(errors.ShowErrorPattern, errors.GetKnowledgeBaseErrMsg)
		return err
	}

	updated, err := e.editTUIKnowledgeBase(cmd.Context(), kb)
	if err != nil {
		e.lg.Zap.Error(errors.EditTUIKnowledgeBaseErrMsg, zap.Error(err))
		fmt.Printf(errors.ShowErrorPattern, errors.EditTUIKnowledgeBaseErrMsg)
		return err
	}

	return e.commitChanges(cmd.Context(), man, updated)
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
		err := fmt.Errorf(errors.ReceivedResponseWasNotExpectedErrMsg)
		e.lg.Zap.Error(err.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}

	if result.IsQuitted || !result.BaseInfo.IsEdited || !result.Description.IsEdited {
		err := fmt.Errorf(errors.KnowledgeBaseWasNotEditedErrMsg)
		e.lg.Zap.Error(err.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}

	if result.BaseInfo.ID == "" {
		err := fmt.Errorf(errors.IDIsEmptyErrMsg)
		e.lg.Zap.Error(err.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	if result.BaseInfo.ShortName == "" {
		err := fmt.Errorf(errors.ShortNameIsEmptyErrMsg)
		e.lg.Zap.Error(err.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}

	var updated kbEnt.KnowledgeBase = previous
	updated.BaseInfo.ID = result.BaseInfo.ID
	updated.BaseInfo.ShortName = result.BaseInfo.ShortName
	updated.BaseInfo.ModifiedDate = result.BaseInfo.ModifiedDate
	updated.ExtraData.Description = result.Description.Description

	return updated, nil
}

func (e Edit) commitChanges(ctx context.Context, man *kbManagement.Management, updated kbEnt.KnowledgeBase) error {
	fmt.Printf("\n---\n👍 you have entered new information about the knowledge base\n")

	err := man.Update(ctx, updated)
	if err != nil {
		e.lg.Zap.Error(errors.UpdateKnowledgeBaseErrMsg, zap.Error(err))
		fmt.Printf(errors.ShowErrorPattern, errors.UpdateKnowledgeBaseErrMsg)
		return err
	}
	fmt.Print("\n---\n👍 changes of the knowledge base have been committed\n")
	return nil
}
