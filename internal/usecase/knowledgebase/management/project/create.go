package project

import (
	"context"

	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"
)

func (p Project) Create(ctx context.Context, name string) error {
	uuid := uuidGen.New().String()
	filePath := p.pm.GetProjectPath(uuid)
	err := p.fc.CreateFile(filePath)
	if err != nil {
		p.lg.Error("fail to create a database file for the project of the knowledge base", zap.Error(err))
		return err
	}

	err = p.man.CreateKnowledgeBase(ctx, uuid, name)
	if err != nil {
		p.lg.Error("fail to insert knowledge base to the project's storage", zap.Error(err))
		return err
	}

	return nil
}
