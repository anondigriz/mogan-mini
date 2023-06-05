package pattern

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Pattern) Get(knowledgeBaseUUID, uuid string) (kbEnt.Pattern, error) {
	k, err := kb.st.GetPattern(knowledgeBaseUUID, uuid)
	if err != nil {
		kb.lg.Error(errMsgs.GetPatternFromStorageFail, zap.Error(err))
		return kbEnt.Pattern{}, errors.WrapStorageFailErr(err)
	}

	return k, nil
}
