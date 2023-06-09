package parameter

import (
	"go.uber.org/zap"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
)

type Parameter struct {
	st *kbSt.Storage
	lg *zap.Logger
}

func New(lg *zap.Logger, st *kbSt.Storage) *Parameter {
	p := &Parameter{
		st: st,
		lg: lg,
	}
	return p
}
