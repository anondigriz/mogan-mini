package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	chooseCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/baseinfo/choose"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
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
	if err := a.showTUIKnowledgeBases(); err != nil {
		a.lg.Zap.Error(errMsgs.ShowTUIKnowledgeBasesFail, zap.Error(err))
		messages.PrintFail(errMsgs.ShowTUIKnowledgeBasesFail)
		return err
	}
	return nil
}

func (a All) showTUIKnowledgeBases() error {
	messages.PrintAllKnowledgeBases()
	kbsu := kbsUC.New(a.lg.Zap,
		kbsSt.New(a.lg.Zap, a.cfg.WorkspaceDir))
	kbs := kbsu.GetAllKnowledgeBases()
	info := make([]kbEnt.BaseInfo, 0, len(kbs))
	for _, kb := range kbs {
		info = append(info, kb.BaseInfo)
	}

	ch := chooseCLI.New(a.lg.Zap)
	ch.ChooseTUI(info)

	return nil
}
