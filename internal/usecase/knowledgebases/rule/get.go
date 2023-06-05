package rule

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Rule) Get(knowledgeBaseUUID, uuid string) (kbEnt.Rule, error) {
	k, err := kb.st.GetRule(knowledgeBaseUUID, uuid)
	if err != nil {
		kb.lg.Error(errMsgs.GetRuleFromStorageFail, zap.Error(err))
		return kbEnt.Rule{}, errors.WrapStorageFailErr(err)
	}

	return k, nil
}
