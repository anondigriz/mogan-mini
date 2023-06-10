package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	choicesTUI "github.com/anondigriz/mogan-mini/internal/tui/choices"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

const (
	removeQuestion string = "Are you sure?"
)

var (
	removeConfirmChoices = []string{"Confirm ✓", "Abort ✕"}
)

type Remove struct {
	lg     *logger.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	kbUUID string
}

func NewRemove(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Remove {
	remove := &Remove{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	remove.Cmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove the knowledge base",
		Long:  `Remove the knowledge base from the working directory`,
		RunE:  remove.runE,
	}
	return remove
}

func (r *Remove) Init() {
	r.Cmd.PersistentFlags().StringVar(&r.kbUUID, "uuid", "", "UUID of the knowledge base")
	cobra.OnInitialize(r.initConfig)
}

func (r *Remove) initConfig() {
}

func (r *Remove) runE(cmd *cobra.Command, args []string) error {
	if r.kbUUID == "" {
		kbsu := kbsUC.New(r.lg.Zap,
			kbsSt.New(r.lg.Zap, r.cfg.WorkspaceDir))
		uuid, err := chooseKnowledgeBase(r.lg.Zap, kbsu)
		if err != nil {
			r.lg.Zap.Error(errMsgs.ChooseKnowledgeBaseFail, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseKnowledgeBaseFail)
			return err
		}
		r.kbUUID = uuid
	}
	messages.PrintKnowledgeBaseChosen(r.kbUUID)

	check, err := r.askTUIConfirm()
	if err != nil {
		r.lg.Zap.Error(errMsgs.AskTUIConfirmFail, zap.Error(err))
		messages.PrintFail(errMsgs.AskTUIConfirmFail)
		return err
	}

	if !check {
		err = fmt.Errorf(errMsgs.NotConfirm)
		r.lg.Zap.Error(errMsgs.NotConfirm)
		messages.PrintFail(errMsgs.NotConfirm)
		return err
	}

	if err = r.remove(cmd.Context()); err != nil {
		r.lg.Zap.Error(errMsgs.RemoveKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.RemoveKnowledgeBaseFail)
		return err
	}
	messages.PrintKnowledgeBaseRemoved(r.kbUUID)

	return r.updateConfig()
}

func (r Remove) askTUIConfirm() (bool, error) {
	messages.PrintConfirmRemoveKnowledgeBase()
	mt := choicesTUI.New(removeQuestion, removeConfirmChoices)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		r.lg.Zap.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return false, err
	}
	result, ok := m.(choicesTUI.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		r.lg.Zap.Error(err.Error())
		return false, err
	}

	if result.IsQuitted {
		err := fmt.Errorf(errMsgs.KnowledgeBaseWasNotChosen)
		r.lg.Zap.Error(err.Error())
		return false, err
	}

	if result.Choice == removeConfirmChoices[0] {
		return true, nil
	}
	return false, nil
}

func (r Remove) remove(ctx context.Context) error {
	st := kbsSt.New(r.lg.Zap, r.cfg.WorkspaceDir)
	kbsu := kbsUC.New(r.lg.Zap, st)
	err := kbsu.RemoveKnowledgeBase(r.kbUUID)
	if err != nil {
		r.lg.Zap.Error(errMsgs.RemoveKnowledgeBaseFail, zap.Error(err))
		return err
	}
	return nil
}

func (r Remove) updateConfig() error {
	if r.cfg.CurrentKnowledgeBase.UUID != r.kbUUID {
		return nil
	}
	r.vp.Set(kbUUIDConfigPath, "")

	if err := r.vp.WriteConfig(); err != nil {
		r.lg.Zap.Error(errMsgs.UpdateConfig, zap.Error(err))
		messages.PrintFail(errMsgs.UpdateConfig)
		return err
	}
	return nil
}
