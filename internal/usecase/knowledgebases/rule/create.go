package rule

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Rule) Create(knowledgeBaseUUID, shortName string) (string, error) {
	uuid := uuidGen.New().String()
	now := time.Now().UTC()
	rule := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         uuid,
			ID:           uuid,
			ShortName:    shortName,
			CreatedDate:  now,
			ModifiedDate: now,
		},

		InputParameters:  []kbEnt.ParameterRule{},
		OutputParameters: []kbEnt.ParameterRule{},
	}

	if err := kb.st.CreateRule(knowledgeBaseUUID, rule); err != nil {
		kb.lg.Error(errMsgs.CreateRuleInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return uuid, nil
}
