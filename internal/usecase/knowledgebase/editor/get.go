package editor

import (
	"context"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (m Editor) Get(ctx context.Context, uuid string) (kbEnt.KnowledgeBase, error) {
	st, err := m.con.GetStorageByProjectUUID(ctx, uuid)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return kbEnt.KnowledgeBase{}, e
	}
	defer st.Shutdown()

	kb := st.GetKnowledgeBase(ctx)

	return kb, nil
}

func (m Editor) GetAll(ctx context.Context) ([]kbEnt.KnowledgeBase, error) {
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

		kb := st.GetKnowledgeBase(ctx)
		kbs = append(kbs, kb)
	}

	return kbs, nil
}
