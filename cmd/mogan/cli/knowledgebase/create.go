package knowledgebase

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

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
		n, err := c.inputName()
		if err != nil {
			fmt.Printf("\n---\nError entering the name of the knowledge base name: %v\n", err)
			return err
		}
		if n == "" {
			err = fmt.Errorf("You did not enter the knowledge base name!")
			fmt.Printf("\n---\n%v\n", err)
			return err
		}
		c.ShortName = n
	}

	fmt.Printf("\n---\nYou entered the knowledge base name: %s\n", c.ShortName)

	man := kbManagement.New(c.lg.Zap, *c.cfg)
	st, err := man.CreateProject(cmd.Context(), c.ShortName)
	if err != nil {
		c.lg.Zap.Error("fail to create database for the project of the knowledge base", zap.Error(err))
		return err
	}
	defer st.Shutdown()
	err = st.Ping(cmd.Context())
	if err != nil {
		c.lg.Zap.Error("fail to ping database for the project of the knowledge base", zap.Error(err))
		return err
	}
	fmt.Printf("\n---\nEverything all right! The project has been created! The knowledge base name: %s\n", c.ShortName)
	return nil
}

func (c *Create) inputName() (string, error) {
	mt := textInputTui.New("What is the name of the knowledge base?", "Awesome knowledge base")
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Zap.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", err
	}
	result, ok := m.(textInputTui.Model)
	if !ok {
		c.lg.Zap.Error("Received a response form that was not expected")
		return "", fmt.Errorf("Received a response form that was not expected")

	}

	if result.IsQuitted {
		return "", fmt.Errorf("Knowledge base name was not entered")
	}

	n := result.TextInput.Value()
	return n, nil
}
