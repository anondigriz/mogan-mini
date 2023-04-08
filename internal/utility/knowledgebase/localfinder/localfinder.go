package localfinder

import (
	"context"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	kbEnt "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-editor-cli/internal/storage/insqlite/knowledgebase"
	"go.uber.org/zap"
)

type LocalFinder struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *LocalFinder {
	s := &LocalFinder{
		lg:  lg,
		cfg: cfg,
	}
	return s
}

func (lf *LocalFinder) FindInProjectsDir(ctx context.Context) []kbEnt.KnowledgeBase {
	var kbs []kbEnt.KnowledgeBase

	paths := lf.find(lf.cfg.ProjectsPath, ".db")

	for i := 0; i < len(paths); i++ {
		kbInfo, err := lf.GetKnowledgeBase(ctx, paths[i])
		if err != nil {
			lf.lg.Error("fail to get knowledge base info", zap.Error(err))
			continue
		}

		kbs = append(kbs, kbInfo)
	}

	return kbs
}

func (lf *LocalFinder) GetKnowledgeBase(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	dsn := fmt.Sprintf("file:%s", filePath)

	st, err := knowledgebase.New(ctx, lf.lg, dsn, lf.cfg.Databases.LogLevel)
	if err != nil {
		lf.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	defer st.Shutdown()

	id := strings.TrimSuffix(filepath.Base(filePath), path.Ext(filePath))
	kb, err := st.GetKnowledgeBase(ctx, id)
	if err != nil {
		lf.lg.Error("fail to get knowledge base info", zap.Error(err))
	}
	kb.Path = filePath
	return kb, nil
}

func (lf *LocalFinder) find(root, ext string) []string {
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
