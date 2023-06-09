package parameter

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Parameter) Update(knowledgeBaseUUID string, ent kbEnt.Parameter) error {
	if err := p.st.UpdateParameter(knowledgeBaseUUID, ent); err != nil {
		p.lg.Error(errMsgs.UpdateParameterInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
