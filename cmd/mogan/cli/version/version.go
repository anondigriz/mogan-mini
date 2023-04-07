package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mogan",
	Long:  `Print the version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mogan Editor CLI v0.1 -- HEAD")
	},
}
