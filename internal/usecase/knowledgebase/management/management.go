package management

import (
	"context"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/finder"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/manager"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/pathmaker"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/project"
	"github.com/anondigriz/mogan-mini/internal/utility/filecreator"
)

type Management struct {
	manager *manager.Manager
	project *project.Project
}

func New(lg *zap.Logger, cfg config.Config) *Management {
	pm := pathmaker.New(cfg)
	f := finder.New(lg, cfg)
	fc := filecreator.New(lg)

	con := connection.New(lg, cfg, pm)
	man := manager.New(lg, con, f)
	p := project.New(lg, cfg, con, pm, man, fc)

	m := &Management{
		manager: man,
		project: p,
	}
	return m
}

func (m Management) CreateProject(ctx context.Context, name string) error {
	return m.project.Create(ctx, name)
}

func (m Management) RemoveProjectByUUID(ctx context.Context, uuid string) error {
	return m.project.RemoveByUUID(ctx, uuid)
}

func (m Management) Get(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return m.manager.Get(ctx, uuid)
}

func (m Management) GetAll(ctx context.Context) ([]kbEnt.KnowledgeBase, error) {
	return m.manager.GetAll(ctx)
}

func (m Management) Update(ctx context.Context, ent kbEnt.KnowledgeBase) error {
	return m.manager.Update(ctx, ent)
}
