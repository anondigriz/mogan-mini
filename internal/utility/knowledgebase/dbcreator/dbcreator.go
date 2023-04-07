package dbcreator

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
	entKB "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-editor-cli/internal/storage/insqlite/knowledgebase"
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
	file := path.Join(d.cfg.Projects, id.String()+".db")
	return file
}

func (d *DBCreator) Create(ctx context.Context, name string, file string) (*knowledgebase.Storage, error) {
	err := d.createDbFile(file)
	if err != nil {
		d.lg.Error("fail to create a database file for the project of the knowledge base", zap.Error(err))
		return nil, err
	}
	dsn := fmt.Sprintf("file:%s", file)

	st, err := knowledgebase.New(ctx, d.lg, dsn, d.cfg.Databases.LogLevel)
	if err != nil {
		d.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	kb := buildKnowledgeBase(file, name)

	err = st.CreateKnowledgeBase(ctx, kb)
	if err != nil {
		d.lg.Error("fail to create knowledge base int the database", zap.Error(err))
		return nil, err
	}

	return st, nil
}

func (d *DBCreator) createDbFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		file, err := os.Create(file)
		if err != nil {
			d.lg.Error("fail to create a database for the project of the knowledge base in directory",
				zap.Error(err), zap.Reflect("file", file))
			return err
		}
		defer file.Close()
	}
	return nil
}

func buildKnowledgeBase(file string, name string) entKB.KnowledgeBase {
	fileName := filepath.Base(file)
	id := strings.TrimSuffix(fileName, path.Ext(fileName))
	now := time.Now().UTC()
	kb := entKB.KnowledgeBase{
		BaseInfo: entKB.BaseInfo{
			ID:           id,
			ShortName:    name,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	return kb
}
