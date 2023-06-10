package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	chooseCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/baseinfo/choose"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

type Choose struct {
	lg     *logger.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	kbUUID string
}

func NewChoose(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Choose {
	c := &Choose{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	c.Cmd = &cobra.Command{
		Use:   "choose",
		Short: "Choose a knowledge base project to work with",
		Long:  `Choose a knowledge base project from the base project directory to be used in the workspace`,
		RunE:  c.runE,
	}
	return c
}

func (c *Choose) Init() {
	c.Cmd.PersistentFlags().StringVar(&c.kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(c.initConfig)
}

func (c *Choose) initConfig() {
}

func (c *Choose) runE(cmd *cobra.Command, args []string) error {
	if c.kbUUID == "" {
		kbsu := kbsUC.New(c.lg.Zap,
			kbsSt.New(c.lg.Zap, c.cfg.WorkspaceDir))
		uuid, err := chooseKnowledgeBase(c.lg.Zap, kbsu)
		if err != nil {
			c.lg.Zap.Error(errMsgs.ChooseKnowledgeBaseFail, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseKnowledgeBaseFail)
			return err
		}
		c.kbUUID = uuid
	}
	return c.commitChoice()
}

func chooseKnowledgeBase(lg *zap.Logger, kbsu *kbsUC.KnowledgeBases) (string, error) {
	kbs := kbsu.GetAllKnowledgeBases()
	info := make([]kbEnt.BaseInfo, 0, len(kbs))
	for _, kb := range kbs {
		info = append(info, kb.BaseInfo)
	}

	ch := chooseCLI.New(lg)
	messages.PrintChooseKnowledgeBase()
	uuid, err := ch.ChooseTUI(info)
	if err != nil {
		lg.Error(errMsgs.ChooseTUIKnowledgeBaseFail, zap.Error(err))
		return "", err
	}
	return uuid, nil
}

func (c Choose) commitChoice() error {
	messages.PrintKnowledgeBaseChosen(c.kbUUID)
	c.vp.Set(kbUUIDConfigPath, c.kbUUID)
	if err := c.vp.WriteConfig(); err != nil {
		c.lg.Zap.Error(errMsgs.UpdateConfig, zap.Error(err))
		messages.PrintFail(errMsgs.UpdateConfig)
		return err
	}
	return nil
}
