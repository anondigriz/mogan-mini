package parameter

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Parameter) Create(knowledgeBaseUUID string, ent kbEnt.Parameter) (string, error) {
	ent.UUID = uuidGen.New().String()
	if err := p.st.CreateParameter(knowledgeBaseUUID, ent); err != nil {
		p.lg.Error(errMsgs.CreateRuleInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return ent.UUID, nil
}
