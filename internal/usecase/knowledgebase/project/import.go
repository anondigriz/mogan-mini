package project

import (
	"context"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"go.uber.org/zap"

	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
)

func (p Project) Import(ctx context.Context, args argsCore.ImportProject) (string, error) {
	pArgs := parser.ParseXMLArgs{
		KnowledgeBaseUUID: "",
		XMLFile:           args.XMLFile,
		FileName:          args.FileName,
	}

	cont, err := p.parser.Parse(pArgs)
	if err != nil {
		p.lg.Error("Fail to parse xml file", zap.Error(err))
		return "", err
	}

	uuid, err := p.Create(ctx, cont.KnowledgeBase.ShortName)
	if err != nil {
		p.lg.Error("fail to create database for the project of the knowledge base", zap.Error(err))
		return "", err
	}
	cont.KnowledgeBase.UUID = uuid

	err = p.editor.Fill(ctx, cont)
	if err != nil {
		p.lg.Error("fail to fill the database of the knowledge base project by the data from the xml file", zap.Error(err))
		return "", err
	}
	return uuid, nil
}
