package connection

import (
	"context"
	"fmt"
	"os"
	"path"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
)

type Connection struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *Connection {
	s := &Connection{
		lg:  lg,
		cfg: cfg,
	}
	return s
}

func (c Connection) GetByUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	filePath := path.Join(c.cfg.ProjectsPath, uuid+".db")
	return c.GetByPath(ctx, filePath)
}

func (c Connection) GetByPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	if _, err := os.Stat(filePath); err != nil {
		c.lg.Error("knowledge base project does not exist", zap.Error(err))
		return nil, err
	}

	dsn := fmt.Sprintf("file:%s", filePath)
	st, err := kbStorage.New(ctx, c.lg, dsn, c.cfg.Databases.LogLevel)
	if err != nil {
		c.lg.Error("fail to connect a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	return st, nil
}
