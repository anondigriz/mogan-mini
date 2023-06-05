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
	ID        string
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
	c.Cmd.PersistentFlags().StringVar(&c.ShortName, "name", "", "short name of the knowledge base")
	c.Cmd.PersistentFlags().StringVar(&c.ID, "id", "", "id of the knowledge base")

	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) runE(cmd *cobra.Command, args []string) error {
	if c.ShortName == "" {
		shortName, err := c.inputTUIShortName()
		if err != nil {
			c.lg.Zap.Error(errMsgs.InputTUIShortNameFail, zap.Error(err))
			messages.PrintFail(errMsgs.InputTUIShortNameFail)
			return err
		}
		c.ShortName = shortName
	}
	messages.PrintEnteredShortNameKnowledgeBase(c.ShortName)

	if c.ID == "" {
		name, err := c.inputTUIID()
		if err != nil {
			c.lg.Zap.Error(errMsgs.InputTUIIDFail, zap.Error(err))
			messages.PrintFail(errMsgs.InputTUIIDFail)
			return err
		}
		c.ID = name
	}
	messages.PrintEnteredIDKnowledgeBase(c.ID)

	return c.createKnowledgeBase(cmd.Context())
}

func (c Create) inputTUIShortName() (string, error) {
	mt := textInputTui.New("What is the short name of the knowledge base?", "Awesome knowledge base")
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Zap.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return "", err
	}

	result, ok := m.(textInputTui.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	shortName := result.TextInput.Value()
	if result.IsQuitted || shortName == "" {
		err = fmt.Errorf(errMsgs.ShortNameIsEmpty)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	return shortName, nil
}

func (c Create) inputTUIID() (string, error) {
	mt := textInputTui.New("What is the ID of the knowledge base?", "00000000-1111-2222-3333-444444444444")
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Zap.Error(errMsgs.RunTUIProgramFail, zap.Error(err))
		return "", err
	}

	result, ok := m.(textInputTui.Model)
	if !ok {
		err = fmt.Errorf(errMsgs.ReceivedResponseWasNotExpected)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	id := result.TextInput.Value()
	if result.IsQuitted || id == "" {
		err = fmt.Errorf(errMsgs.IDIsEmpty)
		c.lg.Zap.Error(err.Error())
		return "", err
	}

	return id, nil
}

func (c Create) createKnowledgeBase(ctx context.Context) error {

	st := kbsSt.New(c.lg.Zap, c.cfg.WorkspaceDir)
	kbsu := kbsUC.New(c.lg.Zap, st)
	now := time.Now().UTC()
	knowledgeBase := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			ID:           c.ID,
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
