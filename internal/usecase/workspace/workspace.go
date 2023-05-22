package workspace

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

const (
	defaultDirName = "mogan"
)

type Workspace struct {
	lg *zap.Logger
}

type CfgFile struct {
	Path string
	Type string
	Name string
}

func New(lg *zap.Logger) *Workspace {
	ws := &Workspace{
		lg: lg,
	}
	return ws
}

func (ws Workspace) InitWorkspaceDir(workspaceDir string) (string, error) {
	if workspaceDir == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		workspaceDir = path.Join(home, defaultDirName)
	}
	err := os.MkdirAll(workspaceDir, os.ModePerm)
	if err != nil {
		ws.lg.Error(errMsgs.MakeWorkspaceDirFail, zap.Error(err))
		return "", errors.NewMakeWorkspaceDirFailErr(err, workspaceDir)
	}
	return workspaceDir, nil
}

func (ws Workspace) SetCfgFile(vp *viper.Viper, workspaceDir string, cfgFile CfgFile) error {
	vp.SetConfigType(cfgFile.Type)
	if cfgFile.Path != "" {
		// Use config file from the flag.
		vp.SetConfigFile(cfgFile.Path)
		return nil
	}

	// Search config in "$HOME/mogan" directory with name "cfg" (without extension).
	if cfgFile.Type != "yaml" && cfgFile.Type != "json" {
		err := errors.NewNotSupportedConfigTypeErr(cfgFile.Type)
		ws.lg.Error(err.Error(), zap.Error(err))
		return err
	}
	cfgFile.Path = path.Join(workspaceDir, cfgFile.Name+"."+cfgFile.Type)
	vp.SetConfigFile(cfgFile.Path)

	return ws.createCfgFile(cfgFile.Path)
}

func (ws Workspace) createCfgFile(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			ws.lg.Error(fmt.Sprintf("fail to create a file on the path %s", filePath),
				zap.Error(err), zap.String("path", filePath))
			return errors.NewCreateCfgFileFailErr(err, filePath)
		}
		defer file.Close()
	}
	return nil
}
