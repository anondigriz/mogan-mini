package initializer

import (
	"fmt"
	"os"
	"path"

	"github.com/anondigriz/mogan-mini/internal/utility/filecreator"
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

func InitProjectsDir(projectsPath string) (string, error) {
	if projectsPath == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		projectsPath = path.Join(home, "mogan")
	}
	err := os.MkdirAll(projectsPath, os.ModePerm)
	if err != nil {
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
