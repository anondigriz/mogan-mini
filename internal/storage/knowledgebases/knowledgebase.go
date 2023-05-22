package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (st Storage) GetKnowledgeBase(uuid string) (kbEnt.KnowledgeBase, error) {
	cont, err := st.GetContainerByUUID(uuid)

	if err != nil {
		st.lg.Error(errMsgs.GetContainerFail, zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	return cont.KnowledgeBase, nil
}

func (st Storage) UpdateKnowledgeBase(ent kbEnt.KnowledgeBase) error {
	cont, err := st.GetContainerByUUID(ent.UUID)

	if err != nil {
		st.lg.Error(errMsgs.GetContainerFail, zap.Error(err))
		return err
	}

	cont.KnowledgeBase = ent

	return st.SaveContainer(cont)
}
