package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Create(knowledgeBase kbEnt.KnowledgeBase) (string, error) {
	knowledgeBase.UUID = uuidGen.New().String()
	cont := &kbEnt.Container{
		KnowledgeBase: knowledgeBase,
		Groups:        make(map[string]kbEnt.Group),
		Parameters:    make(map[string]kbEnt.Parameter),
		Patterns:      make(map[string]kbEnt.Pattern),
		Rules:         make(map[string]kbEnt.Rule),
	}

	if err := kb.st.CreateContainer(cont); err != nil {
		kb.lg.Error(errMsgs.CreateContainerInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return knowledgeBase.UUID, nil
}
