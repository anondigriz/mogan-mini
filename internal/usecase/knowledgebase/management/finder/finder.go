package finder

import (
	"context"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/manager"
)

type settings struct {
	projectsPath string
}

type Finder struct {
	lg       *zap.Logger
	con      *connection.Connection
	settings settings
	man      *manager.Manager
}

func New(lg *zap.Logger, cfg config.Config, con *connection.Connection, man *manager.Manager) *Finder {
	f := &Finder{
		lg:  lg,
		con: con,
		man: man,
		settings: settings{
			projectsPath: cfg.ProjectsPath,
		},
	}
	return f
}

func (f Finder) FindAllProjects(ctx context.Context) []kbEnt.KnowledgeBase {
	var kbs []kbEnt.KnowledgeBase

	paths := f.find(f.settings.projectsPath, ".db")

	for i := 0; i < len(paths); i++ {
		kbInfo, err := f.FindProjectByPath(ctx, paths[i])
		if err != nil {
			f.lg.Error("fail to get knowledge base info", zap.Error(err))
			continue
		}

		kbs = append(kbs, kbInfo)
	}

	return kbs
}

func (f Finder) FindProjectByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	filePath := path.Join(f.settings.projectsPath, uuid+".db")
	return f.FindProjectByPath(ctx, filePath)
}

func (f Finder) FindProjectByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	st, err := f.con.GetStorageByProjectPath(ctx, filePath)
	if err != nil {
		f.lg.Error("fail to open project of the knowledge base", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	defer st.Shutdown()

	kb, err := f.man.Get(ctx, st)
	if err != nil {
		f.lg.Error("fail to get knowledge base info", zap.Error(err))
	}
	kb.Path = filePath
	return kb, nil
}

func (f Finder) find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			f.lg.Error(fmt.Sprintf("fail to walk the directory %s", f.settings.projectsPath), zap.Error(e))
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, p)
		}
		return nil
	})
	return a
}
