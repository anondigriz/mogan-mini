package pathmaker

import (
	"fmt"
	"path"

	"github.com/anondigriz/mogan-mini/internal/config"
)

type settings struct {
	ProjectsPath string
}

type PathMaker struct {
	settings settings
}

func New(cfg config.Config) *PathMaker {
	pm := &PathMaker{
		settings: settings{
			ProjectsPath: cfg.ProjectsPath,
		},
	}
	return pm
}

func (pm PathMaker) GetProjectPath(uuid string) string {
	return path.Join(pm.settings.ProjectsPath, uuid+".db")
}

func (pm PathMaker) GetStorageDSN(filePath string) string {
	return fmt.Sprintf("file:%s", filePath)
}
