package group

import (
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	chooseCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/group/choose"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	chooseTUI "github.com/anondigriz/mogan-mini/internal/tui/group/choose"
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
		chooseTUI.AllowArgs{Group: true, Parameter: false, Rule: false},
		chooseTUI.ShowArgs{Parameter: true, Rule: true})
	if err != nil {
		lg.Error(errMsgs.ChooseTUIGroupFail, zap.Error(err))
		return "", err
	}

	return uuid, nil
}
