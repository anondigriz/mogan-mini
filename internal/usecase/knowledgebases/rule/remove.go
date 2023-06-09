package rule

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (r Rule) Remove(knowledgeBaseUUID, uuid string) error {
	if err := r.st.RemoveRule(knowledgeBaseUUID, uuid); err != nil {
		r.lg.Error(errMsgs.RemoveRuleFromStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
