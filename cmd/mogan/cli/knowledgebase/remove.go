package knowledgebase

import (
	"fmt"
	"os"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Remove struct {
	lg     *zap.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	kbUUID string
}

func NewRemove(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Remove {
	remove := &Remove{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	remove.Cmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove the knowledge base",
		Long:  `Remove the knowledge base`,
		Run:   remove.run,
	}
	return remove
}

func (r *Remove) Init() {
	r.Cmd.PersistentFlags().StringVar(&r.kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(r.initConfig)
}

func (r *Remove) initConfig() {
}

func (r *Remove) run(cmd *cobra.Command, args []string) {
	if r.kbUUID == "" {
		uuid, err := chooseKnowledgeBase(cmd.Context(), r.lg, *r.cfg)
		if err != nil {
			fmt.Printf("\n---\nThere was a problem when choosing a knowledge base: %v\n", err)
			return
		}
		r.kbUUID = uuid
	}

	fmt.Printf("\n---\nOkay, you have selected a knowledge base project with UUID %s\n", r.kbUUID)
	r.vp.Set("CurrentKnowledgeBase.UUID", r.kbUUID)
	err := r.vp.WriteConfig()
	if err != nil {
		r.lg.Error("fail to write config", zap.Error(err))
		os.Exit(1)
	}
}
