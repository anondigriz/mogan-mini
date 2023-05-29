package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
)

func (st Storage) CreateContainer(cont *kbEnt.Container) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, cont.KnowledgeBase.UUID)
	err := cb.WriteContainer(cont)
	if err != nil {
		st.lg.Error(errMsgs.CreateContainerFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetContainer(uuid string) (*kbEnt.Container, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, uuid)
	cont, err := cb.ReadContainer()

	if err != nil {
		st.lg.Error(errMsgs.GetContainerFail, zap.Error(err))
		return nil, err
	}
	return cont, nil
}

func (st Storage) RemoveContainer(uuid string) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, uuid)
	err := cb.RemoveContainer()
	if err != nil {
		st.lg.Error(errMsgs.DeleteContainerFail, zap.Error(err))
		return err
	}
	return nil
}
