package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
)

func (st Storage) CreateParameter(knowledgeBaseUUID string, parameter kbEnt.Parameter) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.WriteParameter(parameter)
	if err != nil {
		st.lg.Error(errMsgs.CreateParameterFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetParameter(knowledgeBaseUUID string, uuid string) (kbEnt.Parameter, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	gr, err := cb.ReadParameter(uuid)
	if err != nil {
		st.lg.Error(errMsgs.GetParameterFail, zap.Error(err))
		return kbEnt.Parameter{}, err
	}
	return gr, nil
}

func (st Storage) UpdateParameter(ent kbEnt.Parameter) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, ent.UUID)
	err := cb.WriteParameter(ent)
	if err != nil {
		st.lg.Error(errMsgs.UpdateParameterFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) RemoveParameter(knowledgeBaseUUID string, uuid string) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.RemoveParameter(uuid)
	if err != nil {
		st.lg.Error(errMsgs.RemoveParameterFail, zap.Error(err))
		return err
	}
	return nil
}
