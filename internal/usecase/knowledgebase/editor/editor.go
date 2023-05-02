package editor

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/finder"
)

type Editor struct {
	lg     *zap.Logger
	con    *connection.Connection
	finder *finder.Finder
}

func New(lg *zap.Logger, con *connection.Connection, finder *finder.Finder) *Editor {
	m := &Editor{
		lg:     lg,
		con:    con,
		finder: finder,
	}
	return m
}
