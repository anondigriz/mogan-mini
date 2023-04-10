package group

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
