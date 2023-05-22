package knowledgebase

import (
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

func (kb KnowledgeBase) Import(args argsCore.ImportKnowledgeBase) (string, error) {
	uuid := uuidGen.New().String()
	pArgs := parser.ParseXMLArgs{
		KnowledgeBaseUUID: uuid,
		XMLFile:           args.XMLFile,
		FileName:          args.FileName,
	}

	cont, err := kb.parser.Parse(pArgs)
	if err != nil {
		kb.lg.Error(errMsgs.ParseKnowledgeBaseFail, zap.Error(err))
		return "", errors.WrapExchangeErr(err)
	}

	if err := kb.st.SaveContainer(&cont); err != nil {
		kb.lg.Error(errMsgs.SaveContainerInStorageFail, zap.Error(err))
		return "", errors.WrapStorageFailErr(err)
	}

	return uuid, nil
}
