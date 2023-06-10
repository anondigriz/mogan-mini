package group

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (g Group) Update(knowledgeBaseUUID string, ent kbEnt.Group) error {
	if err := g.st.UpdateGroup(knowledgeBaseUUID, ent); err != nil {
		g.lg.Error(errMsgs.UpdateGroupInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
