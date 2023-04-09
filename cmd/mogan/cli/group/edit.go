package group

import (
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Edit struct {
	lg  *zap.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewEdit(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Edit {
	edit := &Edit{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	edit.Cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit the group",
		Long:  `Edit the knowledge base group`,
		Run:   edit.run,
	}
	return edit
}

func (e *Edit) Init() {
	cobra.OnInitialize(e.initConfig)
}

func (e *Edit) initConfig() {
}

func (e *Edit) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
