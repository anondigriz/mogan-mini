package group

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Group) Create(knowledgeBaseUUID string, group kbEnt.Group) (string, error) {
	group.UUID = uuidGen.New().String()
	if err := kb.st.CreateGroup(knowledgeBaseUUID, group); err != nil {
		kb.lg.Error(errMsgs.CreateGroupInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return group.UUID, nil
}
