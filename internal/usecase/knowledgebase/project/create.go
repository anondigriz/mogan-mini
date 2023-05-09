package project

import (
	"context"
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"
)

func (p Project) Create(ctx context.Context, name string) (string, error) {
	uuid := uuidGen.New().String()
	filePath := p.pm.MakeProjectPath(uuid)
	err := p.fc.CreateFile(filePath)
	if err != nil {
		p.lg.Error("fail to create a database file for the project of the knowledge base", zap.Error(err))
		return "", err
	}

	now := time.Now().UTC()
	kb := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         uuid,
			ID:           uuid,
			ShortName:    name,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	err = p.strc.CreateStorage(kb, filePath)

	if err != nil {
		p.lg.Error("fail to create knowledge base int the database", zap.Error(err))
		return "", err
	}

	return uuid, nil
}
