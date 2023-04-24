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

type Remove struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewRemove(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Remove {
	remove := &Remove{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	remove.Cmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove the group",
		Long:  `Remove the group from the knowledge base`,
		RunE:  remove.runE,
	}
	return remove
}

func (r *Remove) Init() {
	cobra.OnInitialize(r.initConfig)
}

func (r *Remove) initConfig() {
}

func (r *Remove) runE(cmd *cobra.Command, args []string) error {
	if r.cfg.CurrentKnowledgeBase.UUID == "" {
		err := fmt.Errorf(errors.KnowledgeBaseNotChosenErrMsg)
		r.lg.Zap.Error(err.Error(), zap.Error(err))
		return err
	}
	return nil
}
