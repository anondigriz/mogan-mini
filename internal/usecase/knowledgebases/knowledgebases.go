package knowledgebases

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"

	kbSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	kbUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases/knowledgebase"
)

type KnowledgeBases struct {
	st     *kbSt.Storage
	kb     *kbUC.KnowledgeBase
	lg     *zap.Logger
	parser *parser.Parser
}

func New(lg *zap.Logger, st *kbSt.Storage) *KnowledgeBases {
	kb := kbUC.New(lg, st)
	p := parser.New(lg)
	m := &KnowledgeBases{
		st:     st,
		kb:     kb,
		lg:     lg,
		parser: p,
	}
	return m
}

func (kbs KnowledgeBases) GetAllKnowledgeBases() []kbEnt.KnowledgeBase {
	return kbs.st.GetAllKnowledgeBases()
}
