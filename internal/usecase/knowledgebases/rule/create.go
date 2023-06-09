package rule

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (r Rule) Create(knowledgeBaseUUID string, ent kbEnt.Rule) (string, error) {
	ent.UUID = uuidGen.New().String()
	if err := r.st.CreateRule(knowledgeBaseUUID, ent); err != nil {
		r.lg.Error(errMsgs.CreateRuleInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return ent.UUID, nil
}
