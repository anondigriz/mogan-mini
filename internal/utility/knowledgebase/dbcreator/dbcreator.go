package dbcreator

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/anondigriz/mogan-editor-cli/internal/config"
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

func (d *DBCreator) Create(ctx context.Context, file string) (*knowledgebase.Storage, error) {
	if _, err := os.Stat(file); err != nil {
		file, err := os.Create(file)
		if err != nil {
			d.lg.Error("fail to create a database for the project of the knowledge base in directory",
				zap.Error(err), zap.Reflect("file", file))
			return nil, err
		}
		file.Close()
	}
	dsn := fmt.Sprintf("file:%s", file)

	st, err := knowledgebase.New(ctx, d.lg, dsn, d.cfg.Databases.LogLevel)

	if err != nil {
		d.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	return st, nil
}
