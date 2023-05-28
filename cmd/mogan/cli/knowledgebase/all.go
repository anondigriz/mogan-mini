package knowledgebase

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	listShowTui "github.com/anondigriz/mogan-mini/internal/tui/listshow"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
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
	st := kbsSt.New(a.lg.Zap, a.cfg.WorkspaceDir)
	kbsu := kbsUC.New(a.lg.Zap, st)
	kbs := kbsu.GetAllKnowledgeBases()

	if len(kbs) == 0 {
		messages.PrintNoDataToShow()
		return nil
	}

	err := a.showTUIKnowledgeBases(kbs)
	if err != nil {
		a.lg.Zap.Error(errMsgs.ShowTUIKnowledgeBases, zap.Error(err))
		messages.PrintFail(errMsgs.ShowTUIKnowledgeBases)
		return err
	}
	return nil
}

func (a All) showTUIKnowledgeBases(kbs []kbEnt.KnowledgeBase) error {
	list := make([]string, 0, len(kbs))
	for _, v := range kbs {
		list = append(list, fmt.Sprintf("Id: %s; Short name: %s; UUID: %s", v.ID, v.ShortName, v.UUID))
	}
	mt := listShowTui.New(list)

	if _, err := tea.NewProgram(mt).Run(); err != nil {
		a.lg.Zap.Error(errMsgs.RunTUIProgram, zap.Error(err))
		return err
	}
	return nil
}
