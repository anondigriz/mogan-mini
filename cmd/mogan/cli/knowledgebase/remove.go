package knowledgebase

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	choicesTui "github.com/anondigriz/mogan-mini/internal/tui/choices"
	kbManagement "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management"
)

var confirmChoices = []string{"Confirm ⚠️", "Abort 🚫"}

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
		uuid, err := chooseKnowledgeBase(cmd.Context(), r.lg.Zap, *r.cfg)
		if err != nil {
			fmt.Printf("\n---\nThere was a problem when choosing a knowledge base: %v\n", err)
			return err
		}
		r.kbUUID = uuid
	}

	fmt.Printf("\n---\nOkay, you have chosen a knowledge base project with UUID %s\n", r.kbUUID)

	fmt.Printf("\n---\n")

	check, err := r.askConfirm()
	if err != nil {
		if err != nil {
			fmt.Printf("\n---\nAn error occurred when requesting confirmation: %v\n", err)
			return err
		}
	}
	if !check {
		err = fmt.Errorf("You have not confirmed the removing of the knowledge base projec")
		fmt.Printf("\n---\n%v\n", err)
		return err
	}

	err = r.updateConfig()
	if err != nil {
		fmt.Printf("\n---\nAn error occurred while updating the configuration: %v\n", err)
		return err
	}

	d := kbManagement.New(r.lg.Zap, *r.cfg)
	err = d.RemoveKnowledgeBaseByUUID(cmd.Context(), r.kbUUID)
	if err != nil {
		fmt.Printf("\n---\nSomething went wrong when trying to delete a local knowledge base project: %v\n", err)
		return err
	}

	fmt.Printf("\n---\nThe local knowledge base project was successfully removed\n")
	return nil
}

func (r *Remove) askConfirm() (bool, error) {
	q := "Confirm the removing of the local knowledge base project. This action cannot be undone."
	mt := choicesTui.New(q, confirmChoices)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		r.lg.Zap.Error("Alas, there's been an error: %v", zap.Error(err))
		return false, err
	}
	result, ok := m.(choicesTui.Model)
	if !ok {
		r.lg.Zap.Error("Received a response form that was not expected")
		return false, fmt.Errorf("Received a response form that was not expected")
	}

	if result.IsQuitted {
		return false, fmt.Errorf("Confirmation of deletion was not received")
	}
	if result.Choice == confirmChoices[0] {
		return true, nil
	}
	return false, nil
}

func (r *Remove) updateConfig() error {
	if r.cfg.CurrentKnowledgeBase.UUID != r.kbUUID {
		return nil
	}
	r.vp.Set("CurrentKnowledgeBase.UUID", "")
	err := r.vp.WriteConfig()
	if err != nil {
		r.lg.Zap.Error("fail to write config", zap.Error(err))
		return err
	}
	return nil
}
