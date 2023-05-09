package editor

import (
	"context"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (m Editor) Update(ctx context.Context, ent kbEnt.KnowledgeBase) error {
	st, err := m.con.GetStorageByProjectUUID(ctx, ent.UUID)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return e
	}
	defer st.Shutdown()

	st.UpdateKnowledgeBase(ctx, ent)
	if err != nil {
		e := errors.NewUpdateKnowledgeBaseStorageErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return err
	}

	return nil
}
