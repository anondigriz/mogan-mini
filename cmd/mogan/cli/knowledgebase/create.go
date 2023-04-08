package knowledgebase

import (
	"fmt"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	textInputTui "github.com/anondigriz/mogan-editor-cli/internal/tui/textinput"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/dbcreator"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Create struct {
	lg        *zap.Logger
	vp        *viper.Viper
	cfg       *config.Config
	Cmd       *cobra.Command
	ShortName string
}

func NewCreate(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Create {
	cr := &Create{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	cr.Cmd = &cobra.Command{
		Use:   "init",
		Short: "Create a local knowledge base",
		Long:  `Create a local knowledge base in the base project directory`,
		Run:   cr.run,
	}
	return cr
}

func (c *Create) Init() {
	c.Cmd.PersistentFlags().StringVar(&c.ShortName, "name", "", "short knowledge base name")
	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) run(cmd *cobra.Command, args []string) {
	if c.ShortName == "" {
		n, err := c.inputName()
		if err != nil {
			fmt.Printf("\n---\nError entering the name of the knowledge base name: %v\n", err)
			return
		}
		if n == "" {
			fmt.Printf("\n---\nYou did not enter the knowledge base name!\n")
			return
		}
		c.ShortName = n
	}

	fmt.Printf("\n---\nYou entered the knowledge base name: %s\n", c.ShortName)

	dc := dbcreator.New(c.lg, *c.cfg)
	st, err := dc.Create(cmd.Context(), c.ShortName, dc.GenerateFilePath())
	if err != nil {
		c.lg.Error("fail to create database for the project of the knowledge base", zap.Error(err))
		return
	}
	defer st.Shutdown()
	err = st.Ping(cmd.Context())
	if err != nil {
		c.lg.Error("fail to ping database for the project of the knowledge base", zap.Error(err))
		return
	}
	fmt.Printf("\n---\nEverything all right! The project has been created!: %s\n", c.ShortName)
}

func (c *Create) inputName() (string, error) {
	mt := textInputTui.New("What is the name of the knowledge base?", "Awesome knowledge base")
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		c.lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", err
	}
	if m, ok := m.(textInputTui.Model); ok && m.TextInput.Value() != "" {
		n := m.TextInput.Value()
		return n, nil
	}
	return "", fmt.Errorf("Knowledge base name was not entered")
}
