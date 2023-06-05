package pattern

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Pattern) Update(ent kbEnt.Pattern) error {
	if err := kb.st.UpdatePattern(ent); err != nil {
		kb.lg.Error(errMsgs.UpdatePatternInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
