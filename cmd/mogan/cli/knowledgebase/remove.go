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
	choicesTui "github.com/anondigriz/mogan-mini/internal/tui/choices"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

var removeConfirmChoices = []string{"Confirm ✓", "Abort ✕"}

const (
	removeQuestion string = "Confirm the removing of the local knowledge base project. This action cannot be undone."
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
		Short: "Remove the local knowledge base",
		Long:  `Remove the local knowledge base from the project base directory`,
		RunE:  remove.runE,
	}
	return remove
}

func (r *Remove) Init() {
	r.Cmd.PersistentFlags().StringVar(&r.kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(r.initConfig)
}

func (r *Remove) initConfig() {
}

func (r *Remove) runE(cmd *cobra.Command, args []string) error {
	if r.kbUUID == "" {
		choose := NewChoose(r.lg, r.vp, r.cfg)
		uuid, err := choose.chooseKnowledgeBase(cmd.Context())
		if err != nil {
			r.lg.Zap.Error(errMsgs.ChooseKnowledgeBase, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseKnowledgeBase)
			return err
		}
		r.kbUUID = uuid
	}
	messages.PrintChosenKnowledgeBase(r.kbUUID)

	check, err := r.askTUIConfirm()
	if err != nil {
		r.lg.Zap.Error(errMsgs.AskTUIConfirm, zap.Error(err))
		messages.PrintFail(errMsgs.AskTUIConfirm)
		return err
	}

	if !check {
		err = fmt.Errorf(errMsgs.NotConfirm)
		r.lg.Zap.Error(errMsgs.NotConfirm)
		messages.PrintFail(errMsgs.NotConfirm)
		return err
	}

	if err = r.remove(cmd.Context()); err != nil {
		r.lg.Zap.Error(errMsgs.RemoveKnowledgeBase, zap.Error(err))
		messages.PrintFail(errMsgs.RemoveKnowledgeBase)
		return err
	}
	messages.PrintKnowledgeBaseRemoved(r.kbUUID)

	return r.updateConfig()
}

func (r Remove) askTUIConfirm() (bool, error) {
	mt := choicesTui.New(removeQuestion, removeConfirmChoices)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		r.lg.Zap.Error(errMsgs.RunTUIProgram, zap.Error(err))
		return false, err
	}
	result, ok := m.(choicesTui.Model)
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
		r.lg.Zap.Error(errMsgs.RemoveKnowledgeBase, zap.Error(err))
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
