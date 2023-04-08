package show

import (
	"context"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	"github.com/anondigriz/mogan-editor-cli/internal/storage/insqlite/knowledgebase"
	"go.uber.org/zap"
)

type KnowledgeBase struct {
	Id           string
	Path         string
	Name         string
	CreatedDate  time.Time
	ModifiedDate time.Time
}

type Show struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *Show {
	s := &Show{
		lg:  lg,
		cfg: cfg,
	}
	return s
}

func (s *Show) FindProjects(ctx context.Context) []KnowledgeBase {
	var kbs []KnowledgeBase

	paths := s.find(s.cfg.ProjectsPath, ".db")

	for i := 0; i < len(paths); i++ {
		kbInfo, err := s.GetKnowledgeBase(ctx, paths[i])
		if err != nil {
			s.lg.Error("fail to get knowledge base info", zap.Error(err))
			continue
		}

		kbs = append(kbs, kbInfo)
	}

	return kbs
}

func (s *Show) GetKnowledgeBase(ctx context.Context, filePath string) (KnowledgeBase, error) {
	dsn := fmt.Sprintf("file:%s", filePath)

	st, err := knowledgebase.New(ctx, s.lg, dsn, s.cfg.Databases.LogLevel)
	if err != nil {
		s.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return KnowledgeBase{}, err
	}
	defer st.Shutdown()

	id := strings.TrimSuffix(filepath.Base(filePath), path.Ext(filePath))
	kb, err := st.GetKnowledgeBase(ctx, id)
	if err != nil {
		s.lg.Error("fail to get knowledge base info", zap.Error(err))
	}
	kbInfo := KnowledgeBase{
		Id:           id,
		Path:         filePath,
		Name:         kb.ShortName,
		CreatedDate:  kb.CreatedDate,
		ModifiedDate: kb.ModifiedDate,
	}
	return kbInfo, nil
}

func (s *Show) find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			s.lg.Error(fmt.Sprintf("fail to walk the directory %s", s.cfg.ProjectsPath), zap.Error(e))
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, p)
		}
		return nil
	})
	return a
}
