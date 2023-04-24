package finder

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
)

type Finder struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *Finder {
	s := &Finder{
		lg:  lg,
		cfg: cfg,
	}
	return s
}

func (f Finder) FindAll(ctx context.Context) []kbEnt.KnowledgeBase {
	var kbs []kbEnt.KnowledgeBase

	paths := f.find(f.cfg.ProjectsPath, ".db")

	for i := 0; i < len(paths); i++ {
		kbInfo, err := f.FindByPath(ctx, paths[i])
		if err != nil {
			f.lg.Error("fail to get knowledge base info", zap.Error(err))
			continue
		}

		kbs = append(kbs, kbInfo)
	}

	return kbs
}

func (f Finder) FindByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	filePath := path.Join(f.cfg.ProjectsPath, uuid+".db")
	return f.FindByPath(ctx, filePath)
}

func (f Finder) FindByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	if _, err := os.Stat(filePath); err != nil {
		f.lg.Error("knowledge base project does not exist", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}

	dsn := fmt.Sprintf("file:%s", filePath)
	st, err := kbStorage.New(ctx, f.lg, dsn, f.cfg.Databases.LogLevel)
	if err != nil {
		f.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	defer st.Shutdown()

	kb, err := st.GetKnowledgeBase(ctx)
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
			f.lg.Error(fmt.Sprintf("fail to walk the directory %s", f.cfg.ProjectsPath), zap.Error(e))
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, p)
		}
		return nil
	})
	return a
}
