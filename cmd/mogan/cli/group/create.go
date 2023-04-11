package group

import (
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Create struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewCreate(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Create {
	create := &Create{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	create.Cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a group",
		Long:  `Create a knowledge base group`,
		RunE:  create.runE,
	}
	return create
}

func (c *Create) Init() {
	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) runE(cmd *cobra.Command, args []string) error {
	if c.cfg.CurrentKnowledgeBase.UUID == "" {
		err := errors.KnowledgeBaseNotChosenErr
		c.lg.Zap.Error(err.Error(), zap.Error(err))
		return err
	}
	return nil
}
