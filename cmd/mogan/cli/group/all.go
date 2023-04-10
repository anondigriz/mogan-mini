package group

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Run:   all.run,
	}
	return all
}

func (a *All) Init() {
	cobra.OnInitialize(a.initConfig)
}

func (a *All) initConfig() {
}

func (a *All) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
