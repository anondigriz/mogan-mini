package rule

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Rule) Create(knowledgeBaseUUID string, rule kbEnt.Rule) (string, error) {
	rule.UUID = uuidGen.New().String()
	if err := kb.st.CreateRule(knowledgeBaseUUID, rule); err != nil {
		kb.lg.Error(errMsgs.CreateRuleInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return rule.UUID, nil
}
