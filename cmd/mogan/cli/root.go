package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/anondigriz/mogan-core/pkg/logger"
	"github.com/anondigriz/mogan-editor/internal/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	debug       bool
	lg          *zap.Logger
	vp          *viper.Viper
	cfg         config.Config
	cfgFile     string
	cfgName     string
	cfgType     string
	projectBase string
	rootCmd     = &cobra.Command{
		Use:   "mogan",
		Short: "mogan is an editor of the Multidimensional Open Gnoseological Active Network",
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
	rootCmd.PersistentFlags().StringVar(&projectBase, "projectbase", "", "base project directory (default is \"$HOME/mogan\")")
	rootCmd.PersistentFlags().StringVar(&cfgName, "cfgname", "cfg", "config file name")
	rootCmd.PersistentFlags().StringVar(&cfgType, "cfgtype", "yaml", "config type")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// cfgType
	vp = viper.New()
	vp.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	err := initLogger(vp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = setProjectBase()
	if err != nil {
		lg.Error("fail to set a project base directory", zap.Error(err))
		os.Exit(1)
	}

	err = setCfgFile(vp)
	if err != nil {
		lg.Error("fail to init config file", zap.Error(err))
		os.Exit(1)
	}

	con, err := config.New(lg, vp)
	if err != nil {
		lg.Error("fail to parse config", zap.Error(err))
		os.Exit(1)
	}
	cfg = con
}

func initLogger(vp *viper.Viper) error {
	log, err := logger.New(debug)
	if err != nil {
		return err
	}
	lg = log
	return nil
}

func setProjectBase() error {
	if projectBase == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			lg.Error("fail to define home directory", zap.Error(err))
			return err
		}
		projectBase = path.Join(home, "mogan")
	}
	vp.Set("projectbase", projectBase)
	return nil
}

func setCfgFile(vp *viper.Viper) error {
	vp.SetConfigType(cfgType)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		return nil
	}
	// Search config in "$HOME/mogan" directory with name "cfg" (without extension).
	vp.AddConfigPath(projectBase)
	vp.SetConfigName(cfgName)

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
