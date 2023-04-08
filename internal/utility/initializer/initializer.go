package initializer

import (
	"fmt"
	"os"
	"path"

	"github.com/anondigriz/mogan-core/pkg/logger"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/filecreator"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Initializer struct {
	lg *zap.Logger
}

type CfgFile struct {
	Path string
	Type string
	Name string
}

func New(lg *zap.Logger) *Initializer {
	in := &Initializer{
		lg: lg,
	}
	return in
}

func InitLogger(debug bool) (*zap.Logger, error) {
	lg, err := logger.New(debug)
	if err != nil {
		return nil, err
	}
	return lg, nil
}

func (in *Initializer) InitProjectsDir(projectsPath string) (string, error) {
	if projectsPath == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			in.lg.Error("fail to define home directory", zap.Error(err))
			return "", err
		}
		projectsPath = path.Join(home, "mogan")
	}
	err := os.MkdirAll(projectsPath, os.ModePerm)
	if err != nil {
		in.lg.Error("fail to create directory project base directory", zap.Error(err))
		return "", err
	}
	return projectsPath, nil
}

func (in *Initializer) SetCfgFile(vp *viper.Viper, projectsPath string, file CfgFile) error {
	vp.SetConfigType(file.Type)

	if file.Path != "" {
		// Use config file from the flag.
		vp.SetConfigFile(file.Path)
		return nil
	}
	// Search config in "$HOME/mogan" directory with name "cfg" (without extension).
	if file.Type != "yaml" && file.Type != "json" {
		return fmt.Errorf("not supported config type")
	}
	file.Path = path.Join(projectsPath, file.Name+"."+file.Type)
	vp.SetConfigFile(file.Path)

	fc := filecreator.New(in.lg)
	err := fc.CreateFile(file.Path)
	if err != nil {
		in.lg.Error("fail to create config file", zap.Error(err))
		return err
	}

	return nil
}