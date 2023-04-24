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
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/manager"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/remover"
)

type Management struct {
	creator    *creator.Creator
	finder     *finder.Finder
	remover    *remover.Remover
	connection *connection.Connection
	manager    *manager.Manager
}

func New(lg *zap.Logger, cfg config.Config) *Management {
	r := remover.New(lg, cfg)

	c := creator.New(lg, cfg)
	con := connection.New(lg, cfg)
	man := manager.New(lg)
	f := finder.New(lg, cfg, con, man)

	m := &Management{
		creator:    c,
		remover:    r,
		finder:     f,
		connection: con,
		manager:    man,
	}
	return m
}

func (m Management) RemoveProjectByUUID(ctx context.Context, uuid string) error {
	return m.remover.RemoveProjectByUUID(ctx, uuid)
}

func (m Management) FindAllProjects(ctx context.Context) []kbEnt.KnowledgeBase {
	return m.finder.FindAllProjects(ctx)
}

func (m Management) FindProjectByUUID(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return m.finder.FindProjectByUUID(ctx, uuid)
}

func (m Management) FindProjectByPath(ctx context.Context, filePath string) (kbEnt.KnowledgeBase, error) {
	return m.finder.FindProjectByPath(ctx, filePath)
}

func (m Management) CreateProject(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return m.creator.CreateProject(ctx, filePath)
}

func (m Management) GetStorageByProjectUUID(ctx context.Context, uuid string) (*kbStorage.Storage, error) {
	return m.connection.GetStorageByProjectUUID(ctx, uuid)
}

func (m Management) GetStorageByProjectPath(ctx context.Context, filePath string) (*kbStorage.Storage, error) {
	return m.connection.GetStorageByProjectPath(ctx, filePath)
}
