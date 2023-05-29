package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
)

func (st Storage) CreateGroup(knowledgeBaseUUID string, group kbEnt.Group) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.WriteGroup(group)
	if err != nil {
		st.lg.Error(errMsgs.CreateGroupFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetGroup(knowledgeBaseUUID string, uuid string) (kbEnt.Group, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	gr, err := cb.ReadGroup(uuid)

	if err != nil {
		st.lg.Error(errMsgs.GetContainerFail, zap.Error(err))
		return kbEnt.Group{}, err
	}
	return gr, nil
}

func (st Storage) RemoveGroup(knowledgeBaseUUID string, uuid string) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.RemoveGroup(uuid)
	if err != nil {
		st.lg.Error(errMsgs.DeleteContainerFail, zap.Error(err))
		return err
	}
	return nil
}
