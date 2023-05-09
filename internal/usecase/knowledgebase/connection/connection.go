package connection

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/pathmaker"
)

type Connection struct {
	lg *zap.Logger
	pm *pathmaker.PathMaker
}

func New(lg *zap.Logger, cfg config.Config, pm *pathmaker.PathMaker) *Connection {
	c := &Connection{
		lg: lg,
		pm: pm,
	}
	return c
}

func (c Connection) GetStorageByProjectUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	filePath := c.pm.MakeProjectPath(uuid)
	return c.GetStorage(ctx, filePath, uuid)
}

func (c Connection) GetStorageByProjectPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	uuid := c.pm.GetProjectUUID(filePath)
	return c.GetStorage(ctx, filePath, uuid)
}

func (c Connection) GetStorage(ctx context.Context, filePath string, uuid string) (*kbStorage.Storage, error) {
	if _, err := os.Stat(filePath); err != nil {
		c.lg.Error("knowledge base project does not exist", zap.Error(err))
		return nil, err
	}

	st, err := kbStorage.New(ctx, c.lg, filePath, uuid)
	if err != nil {
		c.lg.Error("fail to connect a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	return st, nil
}
