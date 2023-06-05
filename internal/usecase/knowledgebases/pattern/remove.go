package pattern

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Pattern) Remove(knowledgeBaseUUID, uuid string) error {
	if err := kb.st.RemovePattern(knowledgeBaseUUID, uuid); err != nil {
		kb.lg.Error(errMsgs.RemovePatternFromStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
