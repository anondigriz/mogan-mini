package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	entKB "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/logger"
	chooseTui "github.com/anondigriz/mogan-mini/internal/tui/shared/choose"
	kbUseCase "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase"
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
		uuid, err := c.chooseKnowledgeBase(cmd.Context())
		if err != nil {
			c.lg.Zap.Error(errMsgs.ChooseKnowledgeBase, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseKnowledgeBase)
			return err
		}
		c.kbUUID = uuid
	}

	return c.commitChoice()
}

func (c Choose) chooseKnowledgeBase(ctx context.Context) (string, error) {
	kbu := kbUseCase.New(c.lg.Zap, *c.cfg)
	kbs, err := kbu.GetAll(ctx)
	if err != nil {
		c.lg.Zap.Error(errMsgs.GetAllKnowledgeBases, zap.Error(err))
		return "", err
	}

	kbsInfo := make([]entKB.BaseInfo, 0, len(kbs))
	for _, kb := range kbs {
		kbsInfo = append(kbsInfo, kb.BaseInfo)
	}

	uuid, err := c.chooseTUIKnowledgeBase(kbsInfo)
	if err != nil {
		c.lg.Zap.Error(errMsgs.ChooseTUIKnowledgeBase, zap.Error(err))
		return "", err
	}
	return uuid, nil
}

func (c Choose) chooseTUIKnowledgeBase(kbs []entKB.BaseInfo) (string, error) {
	mt := chooseTui.New(kbs)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Zap.Error(errMsgs.RunTUIProgram, zap.Error(err))
		return "", err
	}
	result, ok := m.(chooseTui.Model)
	if !ok {
		err := fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	if result.IsQuitted {
		err := fmt.Errorf(errMsgs.KnowledgeBaseWasNotChosen)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	return result.ChosenUUID, nil
}

func (c Choose) commitChoice() error {
	messages.PrintChosenKnowledgeBase(c.kbUUID)
	c.vp.Set(kbUUIDConfigPath, c.kbUUID)

	if err := c.vp.WriteConfig(); err != nil {
		c.lg.Zap.Error(errMsgs.UpdateConfig, zap.Error(err))
		messages.PrintFail(errMsgs.UpdateConfig)
		return err
	}
	return nil
}
