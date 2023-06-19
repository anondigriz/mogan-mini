package group

import (
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	chooseCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/group/choose"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

func chooseGroup(lg *zap.Logger, kbsu *kbsUC.KnowledgeBases, knowledgeBaseUUID string) (string, error) {
	ch := chooseCLI.New(lg)
	if err := ch.Init(kbsu, knowledgeBaseUUID); err != nil {
		lg.Error(errMsgs.InitGroupChooserFail, zap.Error(err))
		messages.PrintFail(errMsgs.InitGroupChooserFail)
		return "", err
	}

	messages.PrintChooseGroup()
	uuid, err := ch.ChooseTUI(
		chooseCLI.AllowArgs{Group: true, Parameter: false, Rule: false},
		chooseCLI.ShowArgs{Parameter: true, Rule: true})
	if err != nil {
		lg.Error(errMsgs.ChooseTUIGroupFail, zap.Error(err))
		return "", err
	}

	return uuid, nil
}
