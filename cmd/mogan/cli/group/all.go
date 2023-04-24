package group

import (
	"fmt"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
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
	all := &All{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	all.Cmd = &cobra.Command{
		Use:   "all",
		Short: "Show all groups",
		Long:  `Show all groups in the knowledge base`,
		RunE:  all.runE,
	}
	return all
}

func (a *All) Init() {
	cobra.OnInitialize(a.initConfig)
}

func (a *All) initConfig() {
}

func (a *All) runE(cmd *cobra.Command, args []string) error {
	if a.cfg.CurrentKnowledgeBase.UUID == "" {
		err := fmt.Errorf(errors.KnowledgeBaseNotChosenErrMsg)
		a.lg.Zap.Error(err.Error(), zap.Error(err))
		return err
	}
	return nil
}
