package knowledgebase

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Remove(uuid string) error {
	if err := kb.st.RemoveContainerByUUID(uuid); err != nil {
		kb.lg.Error(errMsgs.SaveContainerInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
