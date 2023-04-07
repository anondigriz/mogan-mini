package create

import (
	"fmt"

	createModel "github.com/anondigriz/mogan-editor-cli/internal/tui/create"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)

var (
	name      string
	CreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a local knowledge base",
		Long:  `Create a local knowledge base in the base project directory`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO run tea

			if name == "" {
				n, err := inputName()
				if err != nil {
					fmt.Printf("\n---\nError entering the name of the knowledge base name: %v\n", err)
					return
				}
				if n == "" {
					fmt.Printf("\n---\nYou did not enter the knowledge base name\n")
					return
				}
				name = n
			}

			fmt.Printf("\n---\nYou entered the knowledge base name: %s\n", name)
		},
	}
)

func init() {
	CreateCmd.PersistentFlags().StringVar(&name, "name", "", "config file")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}

func inputName() (string, error) {
	p := tea.NewProgram(createModel.InitialNameModel())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return "", err
	}
	if m, ok := m.(createModel.NameModel); ok && m.TextInput.Value() != "" {
		n := m.TextInput.Value()
		return n, nil
	}
	return "", fmt.Errorf("knowledge base name was not entered")
}
