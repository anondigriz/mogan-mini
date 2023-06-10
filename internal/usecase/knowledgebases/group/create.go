package group

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (g Group) Create(knowledgeBaseUUID string, ent kbEnt.Group) (string, error) {
	ent.UUID = uuidGen.New().String()
	ent.Groups = make(map[string]kbEnt.Group)
	ent.Parameters = make([]string, 0)
	ent.Rules = make([]string, 0)
	if err := g.st.CreateGroup(knowledgeBaseUUID, ent); err != nil {
		g.lg.Error(errMsgs.CreateGroupInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return ent.UUID, nil
}
