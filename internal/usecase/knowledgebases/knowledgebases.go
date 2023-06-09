package knowledgebases

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	grUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases/group"
	kbUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases/knowledgebase"
	parUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases/parameter"
	patUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases/pattern"
	rulUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases/rule"
)

type KnowledgeBases struct {
	st     *kbSt.Storage
	kb     *kbUC.KnowledgeBase
	gr     *grUC.Group
	par    *parUC.Parameter
	rul    *rulUC.Rule
	pat    *patUC.Pattern
	lg     *zap.Logger
	parser *parser.Parser
}

func New(lg *zap.Logger, st *kbSt.Storage) *KnowledgeBases {
	kb := kbUC.New(lg, st)
	gr := grUC.New(lg, st)
	par := parUC.New(lg, st)
	pat := patUC.New(lg, st)
	rul := rulUC.New(lg, st)
	p := parser.New(lg)
	m := &KnowledgeBases{
		st:     st,
		kb:     kb,
		gr:     gr,
		par:    par,
		pat:    pat,
		rul:    rul,
		lg:     lg,
		parser: p,
	}
	return m
}
