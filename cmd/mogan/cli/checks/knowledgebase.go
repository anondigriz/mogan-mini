package checks

import (
	"fmt"

	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
)

func IsKnowledgeBaseChosen(lg zap.Logger, uuid string) error {
	if uuid == "" {
		err := fmt.Errorf(errMsgs.KnowledgeBaseNotChosen)
		lg.Error(err.Error(), zap.Error(err))
		messages.PrintFail(err.Error())
		return err
	}
	return nil
}
