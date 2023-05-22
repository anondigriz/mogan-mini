package knowledgebase

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Create(name string) (string, error) {
	uuid := uuidGen.New().String()
	now := time.Now().UTC()
	cont := &kbEnt.Container{
		KnowledgeBase: kbEnt.KnowledgeBase{
			BaseInfo: kbEnt.BaseInfo{
				UUID:         uuid,
				ID:           uuid,
				ShortName:    name,
				CreatedDate:  now,
				ModifiedDate: now,
			},
		},
		Groups:     make(map[string]kbEnt.Group),
		Parameters: make(map[string]kbEnt.Parameter),
		Patterns:   make(map[string]kbEnt.Pattern),
		Rules:      make(map[string]kbEnt.Rule),
	}

	if err := kb.st.SaveContainer(cont); err != nil {
		kb.lg.Error(errMsgs.SaveContainerInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return uuid, nil
}
