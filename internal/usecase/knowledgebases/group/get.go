package group

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (g Group) Get(knowledgeBaseUUID, uuid string) (kbEnt.Group, error) {
	group, err := g.st.GetGroup(knowledgeBaseUUID, uuid)
	if err != nil {
		g.lg.Error(errMsgs.GetGroupFromStorageFail, zap.Error(err))
		return kbEnt.Group{}, errors.WrapStorageFailErr(err)
	}

	return group, nil
}

func (g Group) GetAll(knowledgeBaseUUID string) (map[string]kbEnt.Group, error) {
	return g.st.GetAllGroups(knowledgeBaseUUID)
}
