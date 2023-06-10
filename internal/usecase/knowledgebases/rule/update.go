package rule

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (r Rule) Update(knowledgeBaseUUID string, ent kbEnt.Rule) error {
	if err := r.st.UpdateRule(knowledgeBaseUUID, ent); err != nil {
		r.lg.Error(errMsgs.UpdateRuleInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
