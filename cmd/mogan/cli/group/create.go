package group

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Create struct {
	lg  *zap.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewCreate(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Create {
	create := &Create{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	create.Cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a group",
		Long:  `Create a knowledge base group`,
		Run:   create.run,
	}
	return create
}

func (c *Create) Init() {
	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
