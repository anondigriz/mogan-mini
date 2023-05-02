package editor

import (
	"context"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"
)

func (m Editor) Fill(ctx context.Context, cont kbEnt.Container) error {
	st, err := m.con.GetStorageByProjectUUID(ctx, cont.KnowledgeBase.UUID)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return e
	}
	defer st.Shutdown()

	err = st.FillFromContainer(ctx, cont)
	if err != nil {
		m.lg.Error("fail to fill the database of the knowledge base project by the data from the xml file", zap.Error(err))
		return err
	}
	return nil
}
