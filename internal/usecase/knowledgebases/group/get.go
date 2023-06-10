package group

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (g Group) Get(knowledgeBaseUUID, uuid string) (kbEnt.Group, error) {
	groups, err := g.st.GetAllGroups(knowledgeBaseUUID)
	if err != nil {
		g.lg.Error(errMsgs.GetGroupFromStorageFail, zap.Error(err))
		return kbEnt.Group{}, errors.WrapStorageFailErr(err)
	}

	group, ok := g.findInGroups(groups, uuid)
	if !ok {
		err := errors.NewObjectNotFoundErr(uuid)
		g.lg.Error(err.Error(), zap.Error(err))
		return kbEnt.Group{}, err
	}

	return group, nil
}

func (g Group) findInGroups(groups map[string]kbEnt.Group, uuid string) (kbEnt.Group, bool) {
	for _, v := range groups {
		group, ok := g.findInGroup(v, uuid)
		if ok {
			return group, ok
		}
	}
	return kbEnt.Group{}, false
}

func (g Group) findInGroup(group kbEnt.Group, uuid string) (kbEnt.Group, bool) {
	if group.UUID == uuid {
		return group, true
	}
	return g.findInGroups(group.Groups, uuid)

}

func (g Group) GetAll(knowledgeBaseUUID string) (map[string]kbEnt.Group, error) {
	return g.st.GetAllGroups(knowledgeBaseUUID)
}
