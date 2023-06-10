package parameter

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Parameter) Update(knowledgeBaseUUID string, ent kbEnt.Parameter) error {
	if err := p.st.UpdateParameter(knowledgeBaseUUID, ent); err != nil {
		p.lg.Error(errMsgs.UpdateParameterInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
