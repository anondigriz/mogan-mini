package finder

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

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

func (lf *Finder) FindAll(ctx context.Context) []kbEnt.KnowledgeBase {
	var kbs []kbEnt.KnowledgeBase

	paths := lf.find(lf.cfg.ProjectsPath, ".db")

	for i := 0; i < len(paths); i++ {
		kbInfo, err := lf.FindByPath(ctx, paths[i])
		if err != nil {
			lf.lg.Error("fail to get knowledge base info", zap.Error(err))
			continue
		}

		kbs = append(kbs, kbInfo)
	}

	return kbs
}

func (lf *Finder) FindByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	filePath := path.Join(lf.cfg.ProjectsPath, uuid+".db")
	if _, err := os.Stat(filePath); err != nil {
		lf.lg.Error("Knowledge base project does not exist", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err

	}
	return lf.FindByPath(ctx, filePath)
}

func (lf *Finder) FindByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	dsn := fmt.Sprintf("file:%s", filePath)

	st, err := kbStorage.New(ctx, lf.lg, dsn, lf.cfg.Databases.LogLevel)
	if err != nil {
		lf.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	defer st.Shutdown()

	uuid := strings.TrimSuffix(filepath.Base(filePath), path.Ext(filePath))
	kb, err := st.GetKnowledgeBase(ctx)
	if err != nil {
		lf.lg.Error("fail to get knowledge base info", zap.Error(err))
	}
	kb.UUID = uuid
	kb.Path = filePath
	return kb, nil
}

func (lf *Finder) find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			lf.lg.Error(fmt.Sprintf("fail to walk the directory %s", lf.cfg.ProjectsPath), zap.Error(e))
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, p)
		}
		return nil
	})
	return a
}
