package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

func (st Storage) GetKnowledgeBase(uuid string) (kbEnt.KnowledgeBase, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, uuid)
	kb, err := cb.ReadKnowledgeBase(uuid)
	if err != nil {
		st.lg.Error(errMsgs.GetKnowledgeFail, zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	return kb, nil
}

func (st Storage) UpdateKnowledgeBase(ent kbEnt.KnowledgeBase) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, ent.UUID)
	err := cb.WriteKnowledgeBase(ent)
	if err != nil {
		st.lg.Error(errMsgs.UpdateKnowledgeFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetAllKnowledgeBases() map[string]kbEnt.KnowledgeBase {
	fb := filesbroker.New(st.lg, st.KnowledgeBasesDir, "")
	uuids := fb.GetAllChildDirNames()
	result := make(map[string]kbEnt.KnowledgeBase, len(uuids))
	for _, uuid := range uuids {
		kb, err := st.GetKnowledgeBase(uuid)
		if err != nil {
			st.lg.Error(errMsgs.GetKnowledgeFail, zap.Error(err))
			continue
		}
		result[uuid] = kb
	}

	return result
}
