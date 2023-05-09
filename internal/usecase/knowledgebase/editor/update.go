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

	if err := st.Commit(); err != nil {
		m.lg.Error("fail to fill the database of the knowledge base project by the data from the xml file", zap.Error(err))
		return err
	}
	return nil
}
