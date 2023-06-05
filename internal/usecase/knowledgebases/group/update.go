package group

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Group) Update(ent kbEnt.Group) error {
	if err := kb.st.UpdateGroup(ent); err != nil {
		kb.lg.Error(errMsgs.UpdateGroupInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
