package manager

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/finder"
)

type Manager struct {
	lg     *zap.Logger
	con    *connection.Connection
	finder *finder.Finder
}

func New(lg *zap.Logger, con *connection.Connection, finder *finder.Finder) *Manager {
	m := &Manager{
		lg:     lg,
		con:    con,
		finder: finder,
	}
	return m
}
