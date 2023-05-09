package editor

import (
	"context"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (m Editor) Fill(ctx context.Context, cont kbEnt.Container) error {
	st, err := m.con.GetStorageByProjectUUID(ctx, cont.KnowledgeBase.UUID)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return e
	}
	defer st.Shutdown()

	st.FillFromContainer(ctx, cont)
	if err := st.Commit(); err != nil {
		m.lg.Error("fail to fill the database of the knowledge base project by the data from the xml file", zap.Error(err))
		return err
	}
	return nil
}
