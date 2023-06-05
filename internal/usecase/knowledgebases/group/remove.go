package group

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Group) Remove(knowledgeBaseUUID, uuid string) error {
	if err := kb.st.RemoveGroup(knowledgeBaseUUID, uuid); err != nil {
		kb.lg.Error(errMsgs.RemoveGroupFromStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
