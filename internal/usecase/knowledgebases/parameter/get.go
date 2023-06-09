package parameter

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (p Parameter) Get(knowledgeBaseUUID, uuid string) (kbEnt.Parameter, error) {
	parameter, err := p.st.GetParameter(knowledgeBaseUUID, uuid)
	if err != nil {
		p.lg.Error(errMsgs.GetParameterFromStorageFail, zap.Error(err))
		return kbEnt.Parameter{}, errors.WrapStorageFailErr(err)
	}

	return parameter, nil
}

func (p Parameter) GetAll(knowledgeBaseUUID string) (map[string]kbEnt.Parameter, error) {
	return p.st.GetAllParameters(knowledgeBaseUUID)
}
