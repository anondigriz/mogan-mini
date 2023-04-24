package manager

import (
	"context"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (m Manager) Update(ctx context.Context, ent kbEnt.KnowledgeBase) error {
	st, err := m.con.GetStorageByProjectUUID(ctx, ent.UUID)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return e
	}
	defer st.Shutdown()

	err = st.UpdateKnowledgeBase(ctx, ent)
	if err != nil {
		e := errors.NewUpdateKnowledgeBaseStorageErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return err
	}

	return nil
}
