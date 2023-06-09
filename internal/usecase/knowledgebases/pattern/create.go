package pattern

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Pattern) Create(knowledgeBaseUUID string, ent kbEnt.Pattern) (string, error) {
	ent.UUID = uuidGen.New().String()
	if err := p.st.CreatePattern(knowledgeBaseUUID, ent); err != nil {
		p.lg.Error(errMsgs.CreatePatternInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return ent.UUID, nil
}
