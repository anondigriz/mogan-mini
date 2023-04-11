package group

import (
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Edit struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewEdit(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Edit {
	edit := &Edit{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	edit.Cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit the group",
		Long:  `Edit the knowledge base group`,
		RunE:  edit.runE,
	}
	return edit
}

func (e *Edit) Init() {
	cobra.OnInitialize(e.initConfig)
}

func (e *Edit) initConfig() {
}

func (e *Edit) runE(cmd *cobra.Command, args []string) error {
	if e.cfg.CurrentKnowledgeBase.UUID == "" {
		err := errors.KnowledgeBaseNotChosenErr
		e.lg.Zap.Error(err.Error(), zap.Error(err))
		return err
	}
	return nil
}
