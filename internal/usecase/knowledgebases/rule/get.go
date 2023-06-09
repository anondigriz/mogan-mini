package rule

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
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
