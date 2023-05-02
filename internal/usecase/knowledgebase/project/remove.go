package project

import (
	"context"
	"os"

	"go.uber.org/zap"
)

func (p Project) RemoveByUUID(ctx context.Context, uuid string) error {
	filePath := p.pm.MakeProjectPath(uuid)
	if err := os.Remove(filePath); err != nil {
		p.lg.Error("fail to delete the knowledge base project", zap.Error(err))
		return err
	}
	return nil
}
