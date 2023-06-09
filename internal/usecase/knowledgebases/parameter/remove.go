package parameter

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Parameter) Remove(knowledgeBaseUUID, uuid string) error {
	if err := p.st.RemoveParameter(knowledgeBaseUUID, uuid); err != nil {
		p.lg.Error(errMsgs.RemoveParameterFromStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
