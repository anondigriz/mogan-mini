package create

import (
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
		},
	}
)

func init() {
	CreateCmd.PersistentFlags().StringVar(&name, "name", "", "config file")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
