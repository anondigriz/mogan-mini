package remover

import (
	"context"
	"os"
	"path"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
)

type Remover struct {
	lg  *zap.Logger
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *Remover {
	s := &Remover{
		lg:  lg,
		cfg: cfg,
	}
	return s
}

func (lf *Remover) RemoveByUUID(ctx context.Context, uuid string) error {
	filePath := path.Join(lf.cfg.ProjectsPath, uuid+".db")
	if err := os.Remove(filePath); err != nil {
		lf.lg.Error("Fail to delete the knowledge base project", zap.Error(err))
		return err
	}
	return nil
}
