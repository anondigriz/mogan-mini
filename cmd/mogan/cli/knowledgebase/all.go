package knowledgebase

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/logger"
	listShowTui "github.com/anondigriz/mogan-mini/internal/tui/listshow"
	kbUseCase "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase"
)

type All struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewAll(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *All {
	a := &All{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	a.Cmd = &cobra.Command{
		Use:   "all",
		Short: "Show all local knowledge bases",
		Long:  `Show all knowledge base projects that are located in the base project directory`,
		RunE:  a.runE,
	}
	return a
}

func (a *All) Init() {
	cobra.OnInitialize(a.initConfig)
}

func (a *All) initConfig() {
}

func (a *All) runE(cmd *cobra.Command, args []string) error {
	kbu := kbUseCase.New(a.lg.Zap, *a.cfg)
	kbs, err := kbu.GetAll(cmd.Context())
	if err != nil {
		a.lg.Zap.Error(errors.GetAllKnowledgeBasesErrMsg, zap.Error(err))
		messages.PrintFail(errors.GetAllKnowledgeBasesErrMsg)
		return err
	}

	if len(kbs) == 0 {
		messages.PrintNoDataToShow()
		return nil
	}

	err = a.showTUIKnowledgeBases(kbs)
	if err != nil {
		a.lg.Zap.Error(errors.ShowTUIKnowledgeBasesErrMsg, zap.Error(err))
		messages.PrintFail(errors.ShowTUIKnowledgeBasesErrMsg)
		return err
	}
	return nil
}

func (a All) showTUIKnowledgeBases(kbs []kbEnt.KnowledgeBase) error {
	list := make([]string, 0, len(kbs))
	for _, v := range kbs {
		list = append(list, fmt.Sprintf("name: %s; id: %s; path: %s", v.ShortName, v.ID, v.Path))
	}
	mt := listShowTui.New(list)

	if _, err := tea.NewProgram(mt).Run(); err != nil {
		a.lg.Zap.Error(errors.RunTUIProgramErrMsg, zap.Error(err))
		return err
	}
	return nil
}
