package connection

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/loglevel"
	"github.com/anondigriz/mogan-mini/internal/config"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/pathmaker"
)

type settings struct {
	LogLevel loglevel.LogLevel
}

type Connection struct {
	lg       *zap.Logger
	pm       *pathmaker.PathMaker
	settings settings
}

func New(lg *zap.Logger, cfg config.Config, pm *pathmaker.PathMaker) *Connection {
	c := &Connection{
		lg: lg,
		pm: pm,
		settings: settings{
			LogLevel: cfg.Databases.LogLevel,
		},
	}
	return c
}

func (c Connection) GetStorageByProjectUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	filePath := c.pm.MakeProjectPath(uuid)
	return c.GetStorageByProjectPath(ctx, filePath)
}

func (c Connection) GetStorageByProjectPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	if _, err := os.Stat(filePath); err != nil {
		c.lg.Error("knowledge base project does not exist", zap.Error(err))
		return nil, err
	}

	dsn := c.pm.MakeStorageDSN(filePath)
	st, err := kbStorage.New(ctx, c.lg, dsn, c.settings.LogLevel)
	if err != nil {
		c.lg.Error("fail to connect a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	return st, nil
}
