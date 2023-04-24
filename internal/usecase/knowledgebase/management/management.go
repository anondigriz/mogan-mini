package management

import (
	"context"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/creator"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/finder"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/remover"
)

type Management struct {
	creator    *creator.Creator
	finder     *finder.Finder
	remover    *remover.Remover
	connection *connection.Connection
}

func New(lg *zap.Logger, cfg config.Config) *Management {
	r := remover.New(lg, cfg)
	f := finder.New(lg, cfg)
	c := creator.New(lg, cfg)
	con := connection.New(lg, cfg)

	m := &Management{
		creator:    c,
		remover:    r,
		finder:     f,
		connection: con,
	}
	return m
}

func (m Management) RemoveKnowledgeBaseByUUID(ctx context.Context, uuid string) error {
	return m.remover.RemoveByUUID(ctx, uuid)
}

func (m Management) FindAllKnowledgeBase(ctx context.Context) []kbEnt.KnowledgeBase {
	return m.finder.FindAll(ctx)
}

func (m Management) FindKnowledgeBaseByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return m.finder.FindByUUID(ctx, uuid)
}

func (m Management) FindKnowledgeBaseByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	return m.finder.FindByPath(ctx, filePath)
}

func (m Management) CreateKnowledgeBase(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return m.creator.Create(ctx, filePath)
}

func (m Management) GetKnowledgeBaseConnectionByUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	return m.connection.GetByUUID(ctx, uuid)
}

func (m Management) GetKnowledgeBaseConnectionByPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return m.connection.GetByPath(ctx, filePath)
}
