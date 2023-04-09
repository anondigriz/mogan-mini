package connection

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
	"go.uber.org/zap"
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
func (c *Connection) GetByUUID(ctx context.Context, uuid string) (*knowledgebase.Storage, error) {
	filePath := path.Join(c.cfg.ProjectsPath, uuid+".db")
	if _, err := os.Stat(filePath); err != nil {
		c.lg.Error("Knowledge base project does not exist", zap.Error(err))
		return nil, err

	}
	return c.GetByPath(ctx, filePath)
}

func (c *Connection) GetByPath(ctx context.Context, filePath string) (*knowledgebase.Storage, error) {
	dsn := fmt.Sprintf("file:%s", filePath)

	st, err := knowledgebase.New(ctx, c.lg, dsn, c.cfg.Databases.LogLevel)
	if err != nil {
		c.lg.Error("fail to connect a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	return st, nil
}
