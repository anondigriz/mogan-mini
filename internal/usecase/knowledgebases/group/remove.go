package group

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (g Group) Remove(knowledgeBaseUUID, uuid string) error {
	if err := g.st.RemoveGroup(knowledgeBaseUUID, uuid); err != nil {
		g.lg.Error(errMsgs.RemoveGroupFromStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
