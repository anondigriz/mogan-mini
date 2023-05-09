package pathmaker

import (
	"path"
	"path/filepath"
	"strings"

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

func (pm PathMaker) MakeProjectPath(uuid string) string {
	return path.Join(pm.settings.ProjectsPath, uuid+".xml")
}

func (pm PathMaker) GetProjectUUID(filePath string) string {
	return strings.TrimSuffix(filePath, filepath.Ext(filePath))
}
