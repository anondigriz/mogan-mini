package cli

import (
	"fmt"
	"os"

	kbCLI "github.com/anondigriz/mogan-editor-cli/cmd/mogan/cli/knowledgebase"
	"github.com/anondigriz/mogan-editor-cli/internal/config"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/initializer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	isDebug      bool
	isConsoleLog bool
	lg           *zap.Logger
	vp           *viper.Viper
	cfg          *config.Config
	cfgFilePath  string
	cfgFileName  string
	cfgFileType  string
	projectsPath string
	rootCmd      = &cobra.Command{
		Version: "v0.1",
		Use:     "mogan",
		Short:   "mogan is an editor of the Multidimensional Open Gnoseological Active Network",
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
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "enable debug mod")
	rootCmd.PersistentFlags().BoolVarP(&isConsoleLog, "consolelog", "", false, "enable console ")

	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file (default is \"$HOME/mogan/cfg.yaml\")")
	rootCmd.PersistentFlags().StringVar(&projectsPath, "projects", "", "base project directory (default is \"$HOME/mogan\")")
	rootCmd.PersistentFlags().StringVar(&cfgFileName, "cfgname", "cfg", "config file name")
	rootCmd.PersistentFlags().StringVar(&cfgFileType, "cfgtype", "yaml", "config type")

	cobra.OnInitialize(initRootCfg)
	initVars()
	kb := kbCLI.NewRoot(lg, vp, cfg)
	rootCmd.AddCommand(kb.Cmd)
	kb.Init()

	rootCmd.AddCommand(chooseCmd)
}

func initVars() {
	vp = viper.New()
	cfg = &config.Config{}

	log, err := initializer.InitLogger(isDebug)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lg = log
}

func initRootCfg() {
	in := initializer.New(lg)
	newProjectsPath, err := in.InitProjectsDir(projectsPath)
	if err != nil {
		lg.Error("fail to set a project base directory", zap.Error(err))
		os.Exit(1)
	}

	cfgFile := initializer.CfgFile{
		Type: cfgFileType,
		Path: cfgFilePath,
		Name: cfgFileName,
	}

	err = in.SetCfgFile(vp, newProjectsPath, cfgFile)
	if err != nil {
		lg.Error("fail to set config file", zap.Error(err))
		os.Exit(1)
	}

	err = cfg.Fill(lg, vp, isDebug, newProjectsPath)
	if err != nil {
		lg.Error("fail to parse config", zap.Error(err))
		os.Exit(1)
	}

	err = vp.WriteConfig()
	if err != nil {
		lg.Error("fail to write config", zap.Error(err))
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
