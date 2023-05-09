package manager

import (
	"context"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"

	"github.com/anondigriz/mogan-mini/internal/config"
	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/editor"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/finder"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/pathmaker"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/project"
	"github.com/anondigriz/mogan-mini/internal/utility/filecreator"
)

type KnowledgeBase struct {
	editor  *editor.Editor
	project *project.Project
}

func New(lg *zap.Logger, cfg config.Config) *KnowledgeBase {
	pm := pathmaker.New(cfg)
	f := finder.New(lg, cfg)
	fc := filecreator.New(lg)
	strc := knowledgebase.NewStorageCreator(lg)
	con := connection.New(lg, cfg, pm)
	editor := editor.New(lg, con, f)
	parser := parser.New(lg)

	prArgs := project.NewProjectArgs{
		Lg:     lg,
		Cfg:    cfg,
		Con:    con,
		Pm:     pm,
		Editor: editor,
		Fc:     fc,
		Parser: parser,
		Strc:   strc,
	}
	p := project.New(prArgs)

	m := &KnowledgeBase{
		editor:  editor,
		project: p,
	}
	return m
}

func (kb KnowledgeBase) CreateProject(ctx context.Context, name string) (string, error) {
	return kb.project.Create(ctx, name)
}

func (kb KnowledgeBase) ImportProject(ctx context.Context, arg argsCore.ImportProject) (string, error) {
	return kb.project.Import(ctx, arg)
}

func (kb KnowledgeBase) RemoveProjectByUUID(ctx context.Context, uuid string) error {
	return kb.project.RemoveByUUID(ctx, uuid)
}

func (kb KnowledgeBase) Get(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	return kb.editor.Get(ctx, uuid)
}

func (kb KnowledgeBase) GetAll(ctx context.Context) ([]kbEnt.KnowledgeBase, error) {
	return kb.editor.GetAll(ctx)
}

func (kb KnowledgeBase) Update(ctx context.Context, ent kbEnt.KnowledgeBase) error {
	return kb.editor.Update(ctx, ent)
}
