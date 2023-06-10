package pattern

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Pattern) Get(knowledgeBaseUUID, uuid string) (kbEnt.Pattern, error) {
	pattern, err := p.st.GetPattern(knowledgeBaseUUID, uuid)
	if err != nil {
		p.lg.Error(errMsgs.GetPatternFromStorageFail, zap.Error(err))
		return kbEnt.Pattern{}, errors.WrapStorageFailErr(err)
	}

	return pattern, nil
}

func (p Pattern) GetAll(knowledgeBaseUUID string) (map[string]kbEnt.Pattern, error) {
	return p.st.GetAllPatterns(knowledgeBaseUUID)
}
