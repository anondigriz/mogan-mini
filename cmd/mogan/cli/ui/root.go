package ui

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/anondigriz/mogan-mini/ui"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Root struct {
	lg   *logger.Logger
	vp   *viper.Viper
	cfg  *config.Config
	Cmd  *cobra.Command
	addr string
}

func NewRoot(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Root {
	root := &Root{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	root.Cmd = &cobra.Command{
		Use:   "ui",
		Short: "Run UI",
		Long:  `Launch the user interface as a single-page application`,
		Run:   root.run,
	}
	return root
}

func (r *Root) Init() {
	r.Cmd.PersistentFlags().StringVar(&r.addr, "addr", ":4000", "address of UI")
	cobra.OnInitialize(r.initConfig)
}

func (r *Root) initConfig() {
}

func (r *Root) run(cmd *cobra.Command, args []string) {
	// TODO: api := newAPI()
	router := gin.Default()

	// TODO: addRoutes(router, api)

	ui.AddRoutes(router)

	router.Run(r.addr)
}
