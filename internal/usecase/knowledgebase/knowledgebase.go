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

func (k KnowledgeBase) RemoveKnowledgeBaseByUUID(ctx context.Context, uuid string) error {
	return k.management.RemoveKnowledgeBaseByUUID(ctx, uuid)
}

func (k KnowledgeBase) FindAllKnowledgeBase(ctx context.Context) []kbEnt.KnowledgeBase {
	return k.management.FindAllKnowledgeBase(ctx)
}

func (k KnowledgeBase) FindKnowledgeBaseByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return k.management.FindKnowledgeBaseByUUID(ctx, uuid)
}

func (k KnowledgeBase) FindKnowledgeBaseByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	return k.management.FindKnowledgeBaseByUUID(ctx, filePath)
}

func (k KnowledgeBase) CreateKnowledgeBase(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return k.management.CreateKnowledgeBase(ctx, filePath)
}

func (k KnowledgeBase) GetKnowledgeBaseConnectionByUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	return k.management.GetKnowledgeBaseConnectionByUUID(ctx, uuid)
}

func (k KnowledgeBase) GetKnowledgeBaseConnectionByPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return k.management.GetKnowledgeBaseConnectionByPath(ctx, filePath)
}
