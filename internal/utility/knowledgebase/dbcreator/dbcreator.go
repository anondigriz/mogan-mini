package dbcreator

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	entKB "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-editor-cli/internal/storage/insqlite/knowledgebase"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/filecreator"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Migrator interface {
	Migrate() error
}

type DBCreator struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *DBCreator {
	return &DBCreator{
		lg:  lg,
		cfg: cfg,
	}
}

func (d *DBCreator) GenerateFilePath() string {
	id := uuid.New()
	file := path.Join(d.cfg.ProjectsPath, id.String()+".db")
	return file
}

func (d *DBCreator) Create(ctx context.Context, name string, filePath string) (*knowledgebase.Storage, error) {
	fc := filecreator.New(d.lg)
	err := fc.CreateFile(filePath)
	if err != nil {
		d.lg.Error("fail to create a database file for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	dsn := fmt.Sprintf("file:%s", filePath)

	st, err := knowledgebase.New(ctx, d.lg, dsn, d.cfg.Databases.LogLevel)
	if err != nil {
		d.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	kb := buildKnowledgeBase(filePath, name)

	err = st.CreateKnowledgeBase(ctx, kb)
	if err != nil {
		d.lg.Error("fail to create knowledge base int the database", zap.Error(err))
		return nil, err
	}

	return st, nil
}

func buildKnowledgeBase(file string, name string) entKB.KnowledgeBase {
	fileName := filepath.Base(file)
	id := strings.TrimSuffix(fileName, path.Ext(fileName))
	now := time.Now().UTC()
	kb := entKB.KnowledgeBase{
		BaseInfo: entKB.BaseInfo{
			UUID:         id,
			ID:           id,
			ShortName:    name,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	return kb
}
