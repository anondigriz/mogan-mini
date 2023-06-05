package group

import (
	"go.uber.org/zap"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
)

type Group struct {
	st *kbSt.Storage
	lg *zap.Logger
}

func New(lg *zap.Logger, st *kbSt.Storage) *Group {
	m := &Group{
		st: st,
		lg: lg,
	}
	return m
}
