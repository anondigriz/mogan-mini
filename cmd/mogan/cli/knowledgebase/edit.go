package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	editCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/baseinfo/edit"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/checks"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
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
	if err := checks.IsKnowledgeBaseChosen(*e.lg.Zap, e.cfg.CurrentKnowledgeBase.UUID); err != nil {
		return err
	}

	st := kbsSt.New(e.lg.Zap, e.cfg.WorkspaceDir)
	kbsu := kbsUC.New(e.lg.Zap, st)
	kb, err := kbsu.GetKnowledgeBase(e.cfg.CurrentKnowledgeBase.UUID)
	if err != nil {
		e.lg.Zap.Error(errMsgs.GetKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.GetKnowledgeBaseFail)
		return err
	}

	ec := editCLI.New(e.lg.Zap)
	info, err := ec.EditTUI(kb.BaseInfo)
	if err != nil {
		e.lg.Zap.Error(errMsgs.EditTUIKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.EditTUIKnowledgeBaseFail)
		return err
	}
	kb.BaseInfo = info

	return e.commitChanges(kbsu, kb)
}

func (e Edit) commitChanges(kbsu *kbsUC.KnowledgeBases, updated kbEnt.KnowledgeBase) error {
	messages.PrintReceivedNewEntityInfo()

	err := kbsu.UpdateKnowledgeBase(updated)
	if err != nil {
		e.lg.Zap.Error(errMsgs.UpdateKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.UpdateKnowledgeBaseFail)
		return err
	}
	messages.PrintChangesAccepted()
	return nil
}
