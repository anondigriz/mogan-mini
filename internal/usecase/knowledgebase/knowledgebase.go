package knowledgebase

import (
	"context"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management"
)

type KnowledgeBase struct {
	management *management.Management
}

func New(lg *zap.Logger, cfg config.Config) *KnowledgeBase {
	m := management.New(lg, cfg)

	k := &KnowledgeBase{
		management: m,
	}
	return k
}

func (m *KnowledgeBase) RemoveKnowledgeBaseByUUID(ctx context.Context, uuid string) error {
	return m.management.RemoveKnowledgeBaseByUUID(ctx, uuid)
}

func (m *KnowledgeBase) FindAllKnowledgeBase(ctx context.Context) []kbEnt.KnowledgeBase {
	return m.management.FindAllKnowledgeBase(ctx)
}

func (m *KnowledgeBase) FindKnowledgeBaseByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return m.management.FindKnowledgeBaseByUUID(ctx, uuid)
}

func (m *KnowledgeBase) FindKnowledgeBaseByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	return m.management.FindKnowledgeBaseByUUID(ctx, filePath)
}

func (m *KnowledgeBase) CreateKnowledgeBase(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return m.management.CreateKnowledgeBase(ctx, filePath)
}
