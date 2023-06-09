package group

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (g Group) Update(knowledgeBaseUUID string, ent kbEnt.Group) error {
	if err := g.st.UpdateGroup(knowledgeBaseUUID, ent); err != nil {
		g.lg.Error(errMsgs.UpdateGroupInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
