package group

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Root struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewRoot(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Root {
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

	create := NewCreate(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(create.Cmd)
	create.Init()

	edit := NewEdit(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(edit.Cmd)
	edit.Init()

	remove := NewRemove(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(remove.Cmd)
	remove.Init()
}

func (r *Root) initConfig() {
}

func (r *Root) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
