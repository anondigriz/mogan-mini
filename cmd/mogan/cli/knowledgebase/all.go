package knowledgebase

import (
	"fmt"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	kbEnt "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	listshowTui "github.com/anondigriz/mogan-editor-cli/internal/tui/listshow"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/localfinder"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type All struct {
	lg  *zap.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewAll(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *All {
	a := &All{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	a.Cmd = &cobra.Command{
		Use:   "all",
		Short: "Show all local knowledge bases",
		Long:  `Show all knowledge base projects that are located in the base project directory`,
		Run:   a.run,
	}
	return a
}

func (a *All) Init() {
	cobra.OnInitialize(a.initConfig)
}

func (a *All) initConfig() {
}

func (a *All) run(cmd *cobra.Command, args []string) {
	lf := localfinder.New(a.lg, *a.cfg)
	kbs := lf.FindInProjectsDir(cmd.Context())
	err := a.showKnowledgeBases(kbs)
	if err != nil {
		fmt.Printf("\n---\nFail to show list of the knowledge bases: %v\n", err)
		return
	}
}

func (a *All) showKnowledgeBases(kbs []kbEnt.KnowledgeBase) error {
	list := make([]string, 0, len(kbs))
	for _, v := range kbs {
		list = append(list, fmt.Sprintf("name: %s; uuid: %s; path: %s", v.ShortName, v.ID, v.Path))
	}
	mt := listshowTui.New(list)

	if _, err := tea.NewProgram(mt).Run(); err != nil {
		a.lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return err
	}
	return nil
}