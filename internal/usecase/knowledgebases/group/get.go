package group

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Group) Get(knowledgeBaseUUID, uuid string) (kbEnt.Group, error) {
	k, err := kb.st.GetGroup(knowledgeBaseUUID, uuid)
	if err != nil {
		kb.lg.Error(errMsgs.GetGroupFromStorageFail, zap.Error(err))
		return kbEnt.Group{}, errors.WrapStorageFailErr(err)
	}

	return k, nil
}
