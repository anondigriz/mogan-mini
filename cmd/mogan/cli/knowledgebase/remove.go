package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	choicesTui "github.com/anondigriz/mogan-mini/internal/tui/choices"
	kbUseCase "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase"
)

var removeConfirmChoices = []string{"Confirm ✅", "Abort 🚫"}

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
			r.lg.Zap.Error(errors.ChooseKnowledgeBaseErrMsg, zap.Error(err))
			messages.PrintFail(errors.ChooseKnowledgeBaseErrMsg)
			return err
		}
		r.kbUUID = uuid
	}
	messages.PrintChosenKnowledgeBase(r.kbUUID)

	check, err := r.askTUIConfirm()
	if err != nil {
		r.lg.Zap.Error(errors.AskTUIConfirm, zap.Error(err))
		messages.PrintFail(errors.AskTUIConfirm)
		return err
	}

	if !check {
		err = fmt.Errorf(errors.NotConfirmErrMsg)
		r.lg.Zap.Error(errors.NotConfirmErrMsg)
		messages.PrintFail(errors.NotConfirmErrMsg)
		return err
	}

	if err = r.remove(cmd.Context()); err != nil {
		r.lg.Zap.Error(errors.RemoveKnowledgeBaseErrMsg, zap.Error(err))
		messages.PrintFail(errors.RemoveKnowledgeBaseErrMsg)
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
		r.lg.Zap.Error(errors.RunTUIProgramErrMsg, zap.Error(err))
		return false, err
	}
	result, ok := m.(choicesTui.Model)
	if !ok {
		err = fmt.Errorf(errors.ReceivedResponseWasNotExpectedErrMsg)
		r.lg.Zap.Error(err.Error())
		return false, err
	}

	if result.IsQuitted {
		err := fmt.Errorf(errors.KnowledgeBaseWasNotChosenErrMsg)
		r.lg.Zap.Error(err.Error())
		return false, err
	}

	if result.Choice == removeConfirmChoices[0] {
		return true, nil
	}
	return false, nil
}

func (r Remove) remove(ctx context.Context) error {
	kbu := kbUseCase.New(r.lg.Zap, *r.cfg)
	err := kbu.RemoveProjectByUUID(ctx, r.kbUUID)
	if err != nil {
		r.lg.Zap.Error(errors.RemoveKnowledgeBaseErrMsg, zap.Error(err))
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
		r.lg.Zap.Error(errors.UpdateConfigErrMsg, zap.Error(err))
		messages.PrintFail(errors.UpdateConfigErrMsg)
		return err
	}
	return nil
}
