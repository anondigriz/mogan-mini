package group

import (
	"github.com/anondigriz/mogan-editor-cli/internal/config"
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
	all := &All{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	all.Cmd = &cobra.Command{
		Use:   "all",
		Short: "Groups base management",
		Long:  `Show all groups in the knowledge base`,
		Run:   all.run,
	}
	return all
}

func (r *All) Init() {
	cobra.OnInitialize(r.initConfig)
}

func (r *All) initConfig() {
}

func (r *All) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
