package cli

import (
	"fmt"
	"os"

	grCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/group"
	kbCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/knowledgebase"
	uiCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/ui"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	"github.com/anondigriz/mogan-mini/internal/usecase/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	isDebug      bool
	isConsoleLog bool
	lg           *logger.Logger
	vp           *viper.Viper
	cfg          *config.Config
	cfgFilePath  string
	cfgFileName  string
	cfgFileType  string
	workspaceDir string
	rootCmd      = &cobra.Command{
		Version: "v0.1",
		Use:     "mogan",
		Short:   "mogan is an editor of the Multidimensional Open Gnoseological Active Network",
		Long: `A Lightweight and Flexible Editor of the Multidimensional Open Gnoseological Active Network (MOGAN) 
		with love by anondigriz and friends in Go. This is an implementation of a local editor of the MOGAN 
		knowledge bases (aka the Mivar knowledge bases) editor. The MOGAN editor is a mathematical tool for 
		designing artificial intelligence (AI) systems.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "enable debug mod")
	rootCmd.PersistentFlags().BoolVarP(&isConsoleLog, "consolelog", "", false, "enable console ")

	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file (default is \"$HOME/mogan/cfg.yaml\")")
	rootCmd.PersistentFlags().StringVar(&workspaceDir, "workspace", "", "base workspace directory (default is \"$HOME/mogan\")")
	rootCmd.PersistentFlags().StringVar(&cfgFileName, "cfgname", "cfg", "config file name")
	rootCmd.PersistentFlags().StringVar(&cfgFileType, "cfgtype", "yaml", "config type")

	cobra.OnInitialize(initConfig)
	initVars()
	kb := kbCLI.NewRoot(lg, vp, cfg)
	rootCmd.AddCommand(kb.Cmd)
	kb.Init()

	gr := grCLI.NewRoot(lg, vp, cfg)
	rootCmd.AddCommand(gr.Cmd)
	gr.Init()

	u := uiCLI.NewRoot(lg, vp, cfg)
	rootCmd.AddCommand(u.Cmd)
	u.Init()
}

func initVars() {
	vp = viper.New()
	cfg = &config.Config{}

	lg = logger.New()
}

func initConfig() {
	ws := workspace.New(lg.Zap)
	newProjectsPath, err := ws.InitWorkspaceDir(workspaceDir)
	if err != nil {
		lg.Zap.Error("fail to set a project base directory", zap.Error(err))
		os.Exit(1)
	}

	err = lg.Init(newProjectsPath, isDebug)
	if err != nil {
		lg.Zap.Error("fail to init logger", zap.Error(err))
		os.Exit(1)
	}

	cfgFile := workspace.CfgFile{
		Type: cfgFileType,
		Path: cfgFilePath,
		Name: cfgFileName,
	}

	err = ws.SetCfgFile(vp, newProjectsPath, cfgFile)
	if err != nil {
		lg.Zap.Error("fail to set config file", zap.Error(err))
		os.Exit(1)
	}

	err = cfg.Fill(lg.Zap, vp, isDebug, newProjectsPath)
	if err != nil {
		lg.Zap.Error("fail to parse config", zap.Error(err))
		os.Exit(1)
	}

	err = vp.WriteConfig()
	if err != nil {
		lg.Zap.Error("fail to write config", zap.Error(err))
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
