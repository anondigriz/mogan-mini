package knowledgebase

import (
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	textInputTui "github.com/anondigriz/mogan-mini/internal/tui/textinput"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
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
			c.lg.Zap.Error(errMsgs.InputTUIName, zap.Error(err))
			messages.PrintFail(errMsgs.InputTUIName)
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
		c.lg.Zap.Error(errMsgs.RunTUIProgram, zap.Error(err))
		return "", err
	}

	result, ok := m.(textInputTui.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	name := result.TextInput.Value()

	if result.IsQuitted || name == "" {
		err = fmt.Errorf(errMsgs.NameWasNotEntered)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	return name, nil
}

func (c Create) createKnowledgeBase(ctx context.Context) error {
	messages.PrintEnteredShortNameKnowledgeBase(c.ShortName)

	st := kbsSt.New(c.lg.Zap, c.cfg.WorkspaceDir)
	kbsu := kbsUC.New(c.lg.Zap, st)
	now := time.Now().UTC()
	knowledgeBase := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			ShortName:    c.ShortName,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	uuid, err := kbsu.CreateKnowledgeBase(knowledgeBase)
	if err != nil {
		c.lg.Zap.Error(errMsgs.CreateKnowledgeBaseProject, zap.Error(err))
		messages.PrintFail(errMsgs.CreateKnowledgeBaseProject)
		return err
	}
	messages.PrintCreatedKnowledgeBase(uuid)
	return nil
}
