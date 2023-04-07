package usecase

import (
	"github.com/anondigriz/mogan-editor-cli/internal/config"
	"github.com/anondigriz/mogan-editor-cli/internal/core"
	"go.uber.org/zap"
)

type UseCase struct {
	lg  *zap.Logger
	st  *core.Storages
	cfg config.Config
}

func New(cfg config.Config, lg *zap.Logger, st *core.Storages) (*UseCase, error) {
	return &UseCase{
		lg:  lg,
		st:  st,
		cfg: cfg,
	}, nil
}
