package pattern

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Pattern) Create(knowledgeBaseUUID string, pattern kbEnt.Pattern) (string, error) {
	pattern.UUID = uuidGen.New().String()
	if err := kb.st.CreatePattern(knowledgeBaseUUID, pattern); err != nil {
		kb.lg.Error(errMsgs.CreatePatternInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return pattern.UUID, nil
}
