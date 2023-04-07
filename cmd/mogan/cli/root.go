package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/anondigriz/mogan-core/pkg/logger"
	"github.com/anondigriz/mogan-editor-cli/cmd/mogan/cli/create"
	"github.com/anondigriz/mogan-editor-cli/internal/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	debug    bool
	lg       *zap.Logger
	vp       *viper.Viper
	cfg      config.Config
	cfgFile  string
	cfgName  string
	cfgType  string
	projects string
	rootCmd  = &cobra.Command{
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
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mod")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is \"$HOME/mogan/cfg.yaml\")")
	rootCmd.PersistentFlags().StringVar(&projects, "projects", "", "base project directory (default is \"$HOME/mogan\")")
	rootCmd.PersistentFlags().StringVar(&cfgName, "cfgname", "cfg", "config file name")
	rootCmd.PersistentFlags().StringVar(&cfgType, "cfgtype", "yaml", "config type")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(create.CreateCmd)
}

func initConfig() {
	vp = viper.New()
	err := initLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = initProjectsDir()
	if err != nil {
		lg.Error("fail to set a project base directory", zap.Error(err))
		os.Exit(1)
	}

	err = initCfgFile(vp)
	if err != nil {
		lg.Error("fail to init config file", zap.Error(err))
		os.Exit(1)
	}

	con, err := config.New(lg, vp, projects)
	if err != nil {
		lg.Error("fail to parse config", zap.Error(err))
		os.Exit(1)
	}
	cfg = con

	err = vp.WriteConfig()
	if err != nil {
		lg.Error("fail to write config", zap.Error(err))
		os.Exit(1)
	}
}

func initLogger() error {
	log, err := logger.New(debug)
	if err != nil {
		return err
	}
	lg = log
	return nil
}

func initProjectsDir() error {
	if projects == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			lg.Error("fail to define home directory", zap.Error(err))
			return err
		}
		projects = path.Join(home, "mogan")
	}
	err := os.MkdirAll(projects, os.ModePerm)
	if err != nil {
		lg.Error("fail to create directory project base directory", zap.Error(err))
		return err
	}
	return nil
}

func initCfgFile(vp *viper.Viper) error {
	vp.SetConfigType(cfgType)

	if cfgFile != "" {
		// Use config file from the flag.
		vp.SetConfigFile(cfgFile)
		return nil
	}
	// Search config in "$HOME/mogan" directory with name "cfg" (without extension).
	if cfgType != "yaml" && cfgType != "json" {
		return fmt.Errorf("not supported config type")
	}
	cfgFile = path.Join(projects, cfgName+"."+cfgType)
	vp.SetConfigFile(cfgFile)

	_, err := os.Stat(cfgFile)
	if !os.IsExist(err) {
		if _, err := os.Create(cfgFile); err != nil {
			lg.Error("fail to create config file", zap.Error(err))
			return err
		}
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
