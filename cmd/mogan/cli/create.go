package cli

import (
	"fmt"

	textInputTui "github.com/anondigriz/mogan-editor-cli/internal/tui/textinput"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/dbcreator"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var (
	kbName    string
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a local knowledge base",
		Long:  `Create a local knowledge base in the base project directory`,
		Run: func(cmd *cobra.Command, args []string) {
			if kbName == "" {
				n, err := inputName()
				if err != nil {
					fmt.Printf("\n---\nError entering the name of the knowledge base name: %v\n", err)
					return
				}
				if n == "" {
					fmt.Printf("\n---\nYou did not enter the knowledge base name!\n")
					return
				}
				kbName = n
			}

			fmt.Printf("\n---\nYou entered the knowledge base name: %s\n", kbName)

			dc := dbcreator.New(lg, cfg)
			st, err := dc.Create(cmd.Context(), kbName, dc.GenerateFilePath())
			if err != nil {
				lg.Error("fail to create database for the project of the knowledge base", zap.Error(err))
				return
			}
			defer st.Shutdown()
			err = st.Ping(cmd.Context())
			if err != nil {
				lg.Error("fail to ping database for the project of the knowledge base", zap.Error(err))
				return
			}
			fmt.Printf("\n---\nEverything all right! The project has been created!: %s\n", kbName)
		},
	}
)

func initCreateCmd() {
	createCmd.PersistentFlags().StringVar(&kbName, "name", "", "config file")
	cobra.OnInitialize(initCreateCmdCfg)
}

func initCreateCmdCfg() {
}

func inputName() (string, error) {
	mt := textInputTui.New("What is the name of the knowledge base?", "Awesome knowledge base")
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", err
	}
	if m, ok := m.(textInputTui.Model); ok && m.TextInput.Value() != "" {
		n := m.TextInput.Value()
		return n, nil
	}
	return "", fmt.Errorf("Knowledge base name was not entered")
}
