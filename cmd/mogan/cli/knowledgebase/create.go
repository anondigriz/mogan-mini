package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	textInputTui "github.com/anondigriz/mogan-mini/internal/tui/textinput"
	kbManagement "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management"
)

type Create struct {
	lg        *logger.Logger
	vp        *viper.Viper
	cfg       *config.Config
	Cmd       *cobra.Command
	ShortName string
}

func NewCreate(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Create {
	c := &Create{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	c.Cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a local knowledge base",
		Long:  `Create a local knowledge base in the base project directory`,
		RunE:  c.runE,
	}
	return c
}

func (c *Create) Init() {
	c.Cmd.PersistentFlags().StringVar(&c.ShortName, "name", "", "short knowledge base name")
	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) runE(cmd *cobra.Command, args []string) error {
	if c.ShortName == "" {
		name, err := c.inputTUIName()
		if err != nil {
			c.lg.Zap.Error(errors.InputTUINameErrMsg, zap.Error(err))
			fmt.Printf(errors.ShowErrorPattern, errors.InputTUINameErrMsg)
			return err
		}
		c.ShortName = name
	}

	return c.createKnowledgeBase(cmd.Context())
}

func (c Create) inputTUIName() (string, error) {
	mt := textInputTui.New("What is the name of the knowledge base?", "Awesome knowledge base")
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Zap.Error(errors.RunTUIProgramErrMsg, zap.Error(err))
		return "", err
	}

	result, ok := m.(textInputTui.Model)
	if !ok {
		err := fmt.Errorf(errors.ReceivedResponseWasNotExpectedErrMsg)
		c.lg.Zap.Error(err.Error(), zap.Error(err))
		return "", err
	}

	name := result.TextInput.Value()

	if result.IsQuitted || name == "" {
		e := fmt.Errorf(errors.NameWasNotEnteredErrMsg)
		c.lg.Zap.Error(e.Error(), zap.Error(e))
		return "", e
	}

	return name, nil
}

func (c Create) createKnowledgeBase(ctx context.Context) error {
	fmt.Printf("\n---\n👍 you have entered the knowledge base name '%s'\n", c.ShortName)

	man := kbManagement.New(c.lg.Zap, *c.cfg)
	err := man.CreateProject(ctx, c.ShortName)
	if err != nil {
		c.lg.Zap.Error(errors.CreateKnowledgeBaseProjectErrMsg, zap.Error(err))
		fmt.Printf(errors.ShowErrorPattern, errors.CreateKnowledgeBaseProjectErrMsg)
		return err
	}
	return nil
}
