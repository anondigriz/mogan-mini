package knowledgebase

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
		Use:   "kb",
		Short: "Knowledge base management",
		Long:  `Command allows you to manage local knowledge base projects `,
		Run:   root.run,
	}
	return root
}

func (r *Root) Init() {
	cobra.OnInitialize(r.initConfig)

	c := NewCreate(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(c.Cmd)
	c.Init()

	a := NewAll(r.lg, r.vp, r.cfg)
	r.Cmd.AddCommand(a.Cmd)
	a.Init()
}

func (r *Root) initConfig() {
}

func (r *Root) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
