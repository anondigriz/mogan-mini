package create

import (
	"fmt"
	"os"

	createModel "github.com/anondigriz/mogan-editor-cli/internal/tui/create"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)

var (
	name      string
	CreateCmd = &cobra.Command{
		Use:     "create",
		Short:   "Create a local knowledge base",
		Long:    `Create a local knowledge base in the base project directory`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO run tea
			p := tea.NewProgram(createModel.Initial())
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	CreateCmd.PersistentFlags().StringVar(&name, "name", "", "config file")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
