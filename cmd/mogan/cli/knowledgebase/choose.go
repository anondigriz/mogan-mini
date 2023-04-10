package knowledgebase

import (
	"context"
	"fmt"

	entKB "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	chooseTui "github.com/anondigriz/mogan-mini/internal/tui/shared/choose"
	"github.com/anondigriz/mogan-mini/internal/utility/knowledgebase/localfinder"
)

type Choose struct {
	lg     *zap.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	kbUUID string
}

func NewChoose(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Choose {
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
		uuid, err := chooseKnowledgeBase(cmd.Context(), c.lg, *c.cfg)
		if err != nil {
			fmt.Printf("\n---\nThere was a problem when choosing a knowledge base: %v\n", err)
			return err
		}
		c.kbUUID = uuid
	}

	fmt.Printf("\n---\nOkay, you have chosen a knowledge base project with UUID %s\n", c.kbUUID)
	c.vp.Set("CurrentKnowledgeBase.UUID", c.kbUUID)
	err := c.vp.WriteConfig()
	if err != nil {
		fmt.Printf("\n---\nFail to update config %v\n", err)
		c.lg.Error("fail to update config", zap.Error(err))
		return err
	}
	return nil
}

func chooseKnowledgeBase(ctx context.Context, lg *zap.Logger, cfg config.Config) (string, error) {
	lf := localfinder.New(lg, cfg)
	kbs := lf.FindInProjectsDir(ctx)
	bis := make([]entKB.BaseInfo, 0, len(kbs))

	for _, v := range kbs {
		bis = append(bis, v.BaseInfo)
	}

	mt := chooseTui.New(bis)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", err
	}
	result, ok := m.(chooseTui.Model)
	if !ok {
		lg.Error("Received a response form that was not expected")
		return "", fmt.Errorf("Received a response form that was not expected")
	}

	if result.IsQuitted {
		return "", fmt.Errorf("Knowledge base was not chosen")
	}
	return result.ChosenUUID, nil
}
