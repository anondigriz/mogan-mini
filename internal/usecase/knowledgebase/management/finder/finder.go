package finder

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
)

type settings struct {
	ProjectsPath string
}

type Finder struct {
	lg       *zap.Logger
	settings settings
}

func New(lg *zap.Logger, cfg config.Config) *Finder {
	f := &Finder{
		lg: lg,
		settings: settings{
			ProjectsPath: cfg.ProjectsPath,
		},
	}
	return f
}

func (f Finder) FindAllProjects(ctx context.Context) []string {
	var paths []string
	filepath.WalkDir(f.settings.ProjectsPath, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			f.lg.Error(fmt.Sprintf("fail to walk the directory %s", f.settings.ProjectsPath), zap.Error(e))
			return e
		}
		if filepath.Ext(d.Name()) == ".db" {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
