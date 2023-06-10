package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Update(ent kbEnt.KnowledgeBase) error {
	if err := kb.st.UpdateKnowledgeBase(ent); err != nil {
		kb.lg.Error(errMsgs.UpdateKnowledgeBaseInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
