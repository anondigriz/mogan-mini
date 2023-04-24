package manager

import (
	"context"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (m Manager) Get(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	st, err := m.con.GetStorageByProjectUUID(ctx, uuid)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, e
	}
	defer st.Shutdown()

	kb, err := st.GetKnowledgeBase(ctx)
	if err != nil {
		e := errors.NewGetKnowledgeBaseStorageErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, e
	}

	return kb, nil
}

func (m Manager) GetAll(ctx context.Context) ([]kbEnt.KnowledgeBase, error) {
	kbs := []kbEnt.KnowledgeBase{}

	paths := m.finder.FindAllProjects(ctx)

	for _, filePath := range paths {
		st, err := m.con.GetStorageByProjectPath(ctx, filePath)
		if err != nil {
			e := errors.NewPrepareConnectionErr(err)
			m.lg.Error(e.Error(), zap.Error(err))
			continue
		}
		defer st.Shutdown()

		kb, err := st.GetKnowledgeBase(ctx)
		if err != nil {
			e := errors.NewGetKnowledgeBaseStorageErr(err)
			m.lg.Error(e.Error(), zap.Error(err))
			continue
		}
		kb.Path = filePath
		kbs = append(kbs, kb)
	}

	return kbs, nil
}
