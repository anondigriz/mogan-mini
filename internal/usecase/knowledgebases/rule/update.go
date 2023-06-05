package rule

import (
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb Rule) Update(ent kbEnt.Rule) error {
	if err := kb.st.UpdateRule(ent); err != nil {
		kb.lg.Error(errMsgs.UpdateRuleInStorageFail, zap.Error(err))
		return errors.WrapStorageFailErr(err)
	}

	return nil
}
