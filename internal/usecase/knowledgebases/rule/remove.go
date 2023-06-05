package rule

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Rule) Remove(knowledgeBaseUUID, uuid string) error {
	if err := kb.st.RemoveRule(knowledgeBaseUUID, uuid); err != nil {
		kb.lg.Error(errMsgs.RemoveRuleFromStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
