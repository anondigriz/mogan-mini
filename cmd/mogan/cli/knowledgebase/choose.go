package knowledgebase

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	chooseKBTui "github.com/anondigriz/mogan-editor-cli/internal/tui/knowledgebase/choose"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/localfinder"
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
		Run:   c.run,
	}
	return c
}

func (c *Choose) Init() {
	c.Cmd.PersistentFlags().StringVar(&c.kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(c.initConfig)
}

func (c *Choose) initConfig() {
}

func (c *Choose) run(cmd *cobra.Command, args []string) {
	if c.kbUUID == "" {
		uuid, err := c.chooseKnowledgeBase(cmd.Context())
		if err != nil {
			fmt.Printf("\n---\nThere was a problem when choosing a knowledge base: %v\n", err)
			return
		}
		c.kbUUID = uuid
	}

	fmt.Printf("\n---\nOkay, you have selected a knowledge base project with UUID %s\n", c.kbUUID)
	c.vp.Set("CurrentKnowledgeBase.UUID", c.kbUUID)
	err := c.vp.WriteConfig()
	if err != nil {
		c.lg.Error("fail to write config", zap.Error(err))
		os.Exit(1)
	}
}

func (c *Choose) chooseKnowledgeBase(ctx context.Context) (string, error) {
	lf := localfinder.New(c.lg, *c.cfg)
	kbs := lf.FindInProjectsDir(ctx)
	mt := chooseKBTui.New(kbs)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", err
	}
	if m, ok := m.(chooseKBTui.Model); ok && m.Choice != "" {
		return m.Choice, nil
	}
	return "", fmt.Errorf("Knowledge base was not chosen")
}