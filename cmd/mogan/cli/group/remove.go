package group

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Remove struct {
	lg  *zap.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewRemove(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Remove {
	remove := &Remove{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	remove.Cmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove the group",
		Long:  `Remove the group from the knowledge base`,
		Run:   remove.run,
	}
	return remove
}

func (r *Remove) Init() {
	cobra.OnInitialize(r.initConfig)
}

func (r *Remove) initConfig() {
}

func (r *Remove) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
