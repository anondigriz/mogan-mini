package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Get(uuid string) (kbEnt.KnowledgeBase, error) {
	knowledgeBase, err := kb.st.GetKnowledgeBase(uuid)
	if err != nil {
		kb.lg.Error(errMsgs.GetKnowledgeBaseFromStorageFail, zap.Error(err))
		return kbEnt.KnowledgeBase{}, errors.WrapStorageFailErr(err)
	}

	return knowledgeBase, nil
}

func (kb KnowledgeBase) GetAll() map[string]kbEnt.KnowledgeBase {
	return kb.st.GetAllKnowledgeBases()
}
