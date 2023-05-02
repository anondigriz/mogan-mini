package knowledgebase

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

const (
	kbUUIDConfigPath string = "CurrentKnowledgeBase.UUID"
)

func NewRoot(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Root {
	root := &Root{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	root.Cmd = &cobra.Command{
		Use:   "kb",
		Short: "Knowledge base management",
		Long:  `Local knowledge bases projects management`,
		Run:   root.run,
	}
	return root
}

func (r *Root) Init() {
	cobra.OnInitialize(r.initConfig)

	create := NewCreate(r.lg, r.cfg)
	r.Cmd.AddCommand(create.Cmd)
	create.Init()

	all := NewAll(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(all.Cmd)
	all.Init()

	choose := NewChoose(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(choose.Cmd)
	choose.Init()

	edit := NewEdit(r.lg, r.cfg)
	r.Cmd.AddCommand(edit.Cmd)
	edit.Init()

	remove := NewRemove(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(remove.Cmd)
	remove.Init()

	im := NewImport(r.lg, r.cfg)
	r.Cmd.AddCommand(im.Cmd)
	im.Init()
}

func (r *Root) initConfig() {
}

func (r *Root) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
