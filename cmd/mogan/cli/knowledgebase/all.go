package knowledgebase

import (
	"fmt"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/logger"
	listshowTui "github.com/anondigriz/mogan-mini/internal/tui/listshow"
	"github.com/anondigriz/mogan-mini/internal/utility/knowledgebase/localfinder"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	lf := localfinder.New(a.lg.Zap, *a.cfg)
	kbs := lf.FindInProjectsDir(cmd.Context())
	err := a.showKnowledgeBases(kbs)
	if err != nil {
		fmt.Printf("\n---\nFail to show list of the knowledge bases: %v\n", err)
		return err
	}
	return nil
}

func (a *All) showKnowledgeBases(kbs []kbEnt.KnowledgeBase) error {
	list := make([]string, 0, len(kbs))
	for _, v := range kbs {
		list = append(list, fmt.Sprintf("name: %s; uuid: %s; path: %s", v.ShortName, v.ID, v.Path))
	}
	mt := listshowTui.New(list)

	if _, err := tea.NewProgram(mt).Run(); err != nil {
		a.lg.Zap.Error("Alas, there's been an error: %v", zap.Error(err))
		return err
	}
	return nil
}
