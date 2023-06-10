package knowledgebase

import (
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"go.uber.org/zap"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
)

type KnowledgeBase struct {
	st     *kbSt.Storage
	lg     *zap.Logger
	parser *parser.Parser
}

func New(lg *zap.Logger, st *kbSt.Storage) *KnowledgeBase {
	p := parser.New(lg)
	m := &KnowledgeBase{
		st:     st,
		lg:     lg,
		parser: p,
	}
	return m
}
