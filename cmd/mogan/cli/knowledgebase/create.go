package knowledgebase

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	textInputTui "github.com/anondigriz/mogan-mini/internal/tui/textinput"
	kbUseCase "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase"
)

type Create struct {
	lg        *logger.Logger
	cfg       *config.Config
	Cmd       *cobra.Command
	ShortName string
}

func NewCreate(lg *logger.Logger, cfg *config.Config) *Create {
	c := &Create{
		lg:  lg,
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
			messages.PrintFail(errors.InputTUINameErrMsg)
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
		err = fmt.Errorf(errors.ReceivedResponseWasNotExpectedErrMsg)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	name := result.TextInput.Value()

	if result.IsQuitted || name == "" {
		err = fmt.Errorf(errors.NameWasNotEnteredErrMsg)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	return name, nil
}

func (c Create) createKnowledgeBase(ctx context.Context) error {
	messages.PrintEnteredShortNameKnowledgeBase(c.ShortName)

	kbu := kbUseCase.New(c.lg.Zap, *c.cfg)
	uuid, err := kbu.CreateProject(ctx, c.ShortName)
	if err != nil {
		c.lg.Zap.Error(errors.CreateKnowledgeBaseProjectErrMsg, zap.Error(err))
		messages.PrintFail(errors.CreateKnowledgeBaseProjectErrMsg)
		return err
	}
	messages.PrintCreatedKnowledgeBase(uuid)
	return nil
}
