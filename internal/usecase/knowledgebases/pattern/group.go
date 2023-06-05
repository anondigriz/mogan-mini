package pattern

import (
	"go.uber.org/zap"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
)

type Pattern struct {
	st *kbSt.Storage
	lg *zap.Logger
}

func New(lg *zap.Logger, st *kbSt.Storage) *Pattern {
	m := &Pattern{
		st: st,
		lg: lg,
	}
	return m
}
