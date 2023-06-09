package pattern

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Pattern) Update(knowledgeBaseUUID string, ent kbEnt.Pattern) error {
	if err := p.st.UpdatePattern(knowledgeBaseUUID, ent); err != nil {
		p.lg.Error(errMsgs.UpdatePatternInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
