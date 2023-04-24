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
	lg         *zap.Logger
	management *management.Management
}

func New(lg *zap.Logger, cfg config.Config) *KnowledgeBase {
	m := management.New(lg, cfg)

	k := &KnowledgeBase{
		lg:         lg,
		management: m,
	}
	return k
}

func (k KnowledgeBase) RemoveProjectByUUID(ctx context.Context, uuid string) error {
	return k.management.RemoveProjectByUUID(ctx, uuid)
}

func (k KnowledgeBase) FindAllProjects(ctx context.Context) []kbEnt.KnowledgeBase {
	return k.management.FindAllProjects(ctx)
}

func (k KnowledgeBase) FindProjectByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return k.management.FindProjectByUUID(ctx, uuid)
}

func (k KnowledgeBase) FindProjectByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	return k.management.FindProjectByPath(ctx, filePath)
}

func (k KnowledgeBase) CreateProject(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return k.management.CreateProject(ctx, filePath)
}

func (k KnowledgeBase) GetStorageByProjectUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	return k.management.GetStorageByProjectUUID(ctx, uuid)
}

func (k KnowledgeBase) GetStorageByProjectPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return k.management.GetStorageByProjectPath(ctx, filePath)
}
