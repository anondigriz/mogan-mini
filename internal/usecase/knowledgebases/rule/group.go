package rule

import (
	"go.uber.org/zap"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
)

type Rule struct {
	st *kbSt.Storage
	lg *zap.Logger
}

func New(lg *zap.Logger, st *kbSt.Storage) *Rule {
	m := &Rule{
		st: st,
		lg: lg,
	}
	return m
}
