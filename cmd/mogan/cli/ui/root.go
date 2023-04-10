package ui

import (
	"fmt"
	"time"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/anondigriz/mogan-mini/ui"
	ginzap "github.com/gin-contrib/zap"
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
	if !r.cfg.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(r.lg.Zap, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(r.lg.Zap, true))
	// TODO: addRoutes(router, api)

	ui.AddRoutes(router)
	fmt.Printf("UI is launched at: %s\n", r.addr)

	router.Run(r.addr)
}
