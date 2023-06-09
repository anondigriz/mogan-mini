package rule

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (r Rule) Update(knowledgeBaseUUID string, ent kbEnt.Rule) error {
	if err := r.st.UpdateRule(knowledgeBaseUUID, ent); err != nil {
		r.lg.Error(errMsgs.UpdateRuleInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
