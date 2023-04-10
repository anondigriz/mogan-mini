package dbremover

import (
	"context"
	"os"
	"path"

	"github.com/anondigriz/mogan-mini/internal/config"
	"go.uber.org/zap"
)

type DBRemover struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *DBRemover {
	s := &DBRemover{
		lg:  lg,
		cfg: cfg,
	}
	return s
}

func (lf *DBRemover) RemoveByUUID(ctx context.Context, uuid string) error {
	filePath := path.Join(lf.cfg.ProjectsPath, uuid+".db")
	if err := os.Remove(filePath); err != nil {
		lf.lg.Error("Fail to delete the knowledge base project", zap.Error(err))
		return err

	}
	return nil
}
