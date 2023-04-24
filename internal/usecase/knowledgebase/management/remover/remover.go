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

func (r Remover) RemoveProjectByUUID(ctx context.Context, uuid string) error {
	filePath := path.Join(r.cfg.ProjectsPath, uuid+".db")
	if err := os.Remove(filePath); err != nil {
		r.lg.Error("fail to delete the knowledge base project", zap.Error(err))
		return err
	}
	return nil
}
