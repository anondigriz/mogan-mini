package usecase

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/core"
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
