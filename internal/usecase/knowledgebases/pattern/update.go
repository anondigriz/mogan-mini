package pattern

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Pattern) Update(knowledgeBaseUUID string, ent kbEnt.Pattern) error {
	if err := p.st.UpdatePattern(knowledgeBaseUUID, ent); err != nil {
		p.lg.Error(errMsgs.UpdatePatternInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
