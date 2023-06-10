package rule

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (r Rule) Get(knowledgeBaseUUID, uuid string) (kbEnt.Rule, error) {
	rule, err := r.st.GetRule(knowledgeBaseUUID, uuid)
	if err != nil {
		r.lg.Error(errMsgs.GetRuleFromStorageFail, zap.Error(err))
		return kbEnt.Rule{}, errors.WrapStorageFailErr(err)
	}

	return rule, nil
}

func (r Rule) GetAll(knowledgeBaseUUID string) (map[string]kbEnt.Rule, error) {
	return r.st.GetAllRules(knowledgeBaseUUID)
}
