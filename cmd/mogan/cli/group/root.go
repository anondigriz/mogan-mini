package group

import (
	"github.com/anondigriz/mogan-editor-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Root struct {
	lg  *zap.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewRoot(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Root {
	root := &Root{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	root.Cmd = &cobra.Command{
		Use:   "gr",
		Short: "Groups management",
		Long:  `Knowledge base groups management`,
		Run:   root.run,
	}
	return root
}

func (r *Root) Init() {
	cobra.OnInitialize(r.initConfig)

	all := NewAll(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(all.Cmd)
	all.Init()
}

func (r *Root) initConfig() {
}

func (r *Root) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
