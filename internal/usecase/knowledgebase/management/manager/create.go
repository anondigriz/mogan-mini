package manager

import (
	"context"
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (m Manager) CreateKnowledgeBase(ctx context.Context, uuid string, name string) error {
	st, err := m.con.GetStorageByProjectUUID(ctx, uuid)
	if err != nil {
		e := errors.NewPrepareConnectionErr(err)
		m.lg.Error(e.Error(), zap.Error(err))
		return e
	}
	defer st.Shutdown()

	now := time.Now().UTC()
	kb := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         uuid,
			ID:           uuid,
			ShortName:    name,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}

	err = st.CreateKnowledgeBase(ctx, kb)
	if err != nil {
		m.lg.Error("fail to create knowledge base int the database", zap.Error(err))
		return err
	}

	return nil
}
