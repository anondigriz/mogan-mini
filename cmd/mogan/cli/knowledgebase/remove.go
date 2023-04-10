package knowledgebase

import (
	"fmt"

	choicesTui "github.com/anondigriz/mogan-mini/internal/tui/choices"
	"github.com/anondigriz/mogan-mini/internal/utility/knowledgebase/dbremover"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var confirmChoices = []string{"Confirm ⚠️", "Abort 🚫"}

type Remove struct {
	lg     *zap.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	kbUUID string
}

func NewRemove(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Remove {
	remove := &Remove{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	remove.Cmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove the knowledge base",
		Long:  `Remove the knowledge base`,
		Run:   remove.run,
	}
	return remove
}

func (r *Remove) Init() {
	r.Cmd.PersistentFlags().StringVar(&r.kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(r.initConfig)
}

func (r *Remove) initConfig() {
}

func (r *Remove) run(cmd *cobra.Command, args []string) {
	if r.kbUUID == "" {
		uuid, err := chooseKnowledgeBase(cmd.Context(), r.lg, *r.cfg)
		if err != nil {
			fmt.Printf("\n---\nThere was a problem when choosing a knowledge base: %v\n", err)
			return
		}
		r.kbUUID = uuid
	}

	fmt.Printf("\n---\nOkay, you have chosen a knowledge base project with UUID %s\n", r.kbUUID)

	fmt.Printf("\n---\n")

	check, err := r.askConfirm()
	if err != nil {
		if err != nil {
			fmt.Printf("\n---\nAn error occurred when requesting confirmation: %v\n", err)
			return
		}
	}
	if !check {
		fmt.Printf("\n---\nYou have not confirmed the removing of the knowledge base project\n")
		return
	}

	err = r.updateConfig()
	if err != nil {
		fmt.Printf("\n---\nAn error occurred while updating the configuration: %v\n", err)
		return
	}

	d := dbremover.New(r.lg, *r.cfg)
	err = d.RemoveByUUID(cmd.Context(), r.kbUUID)
	if err != nil {
		fmt.Printf("\n---\nSomething went wrong when trying to delete a local knowledge base project: %v\n", err)
		return
	}

	fmt.Printf("\n---\nThe local knowledge base project was successfully removed\n")
}

func (r *Remove) askConfirm() (bool, error) {
	q := "Confirm the removing of the local knowledge base project. This action cannot be undone."
	mt := choicesTui.New(q, confirmChoices)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		r.lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return false, err
	}
	result, ok := m.(choicesTui.Model)
	if !ok {
		r.lg.Error("Received a response form that was not expected")
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
		r.lg.Error("fail to write config", zap.Error(err))
		return err
	}
	return nil
}
