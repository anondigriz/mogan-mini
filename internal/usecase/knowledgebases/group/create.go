package group

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Group) Create(knowledgeBaseUUID, name string) (string, error) {
	uuid := uuidGen.New().String()
	now := time.Now().UTC()
	group := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         uuid,
			ID:           uuid,
			ShortName:    name,
			CreatedDate:  now,
			ModifiedDate: now,
		},
		Groups:     make(map[string]kbEnt.Group),
		Parameters: []string{},
		Rules:      []string{},
	}

	if err := kb.st.CreateGroup(knowledgeBaseUUID, group); err != nil {
		kb.lg.Error(errMsgs.CreateGroupInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return uuid, nil
}
