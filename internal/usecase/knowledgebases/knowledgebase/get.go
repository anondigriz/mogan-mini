package knowledgebase

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Get(uuid string) (kbEnt.KnowledgeBase, error) {
	k, err := kb.st.GetKnowledgeBase(uuid)
	if err != nil {
		kb.lg.Error(errMsgs.GetKnowledgeBaseFromStorageFail, zap.Error(err))
		return kbEnt.KnowledgeBase{}, errors.WrapStorageFailErr(err)
	}

	return k, nil
}
