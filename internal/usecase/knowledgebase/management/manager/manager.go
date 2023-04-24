package manager

import (
	"context"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
)

type Manager struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *Manager {
	m := &Manager{
		lg: lg,
	}
	return m
}

func (m Manager) Get(ctx context.Context, st *kbStorage.Storage) (kbEnt.KnowledgeBase, error) {
	kb, err := st.GetKnowledgeBase(ctx)
	if err != nil {
		m.lg.Error("fail to get knowledge base info", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	return kb, nil
}

func (m Manager) Update(ctx context.Context, st *kbStorage.Storage, ent kbEnt.KnowledgeBase) error {
	err := st.UpdateKnowledgeBase(ctx, ent)
	if err != nil {
		m.lg.Error("error occurred while updating the knowledge base", zap.Error(err))
		return err
	}
	return nil
}
