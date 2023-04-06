package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/anondigriz/mogan-editor/internal/config"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	cfg     config.Config
	rootCmd = &cobra.Command{
		Use:   "mogan",
		Short: "mogan is an Editor of the Multidimensional Open Gnoseological Active Network",
		Long: `A Lightweight and Flexible Editor of the Multidimensional Open Gnoseological Active Network (MOGAN) with
	love by anondigriz and friends in Go. The MOGAN editor is a mathematical 
	tool for designing artificial intelligence (AI) systems. The MOGAN is a 
	combination of the production rule system and Petri nets. The knowledge 
	bases based on the MOGAN are used for semantic analysis and adequate 
	representation of humanitarian epistemological and axiological 
	principles in the process of developing AI.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)




func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mod")
}

func initConfig() {
	cli := config.CLICfg{
		Debug: debug,
	}
	con, err := config.New(cli)
	if err != nil {
		log.Fatal(err)
	}
	cfg = con
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
