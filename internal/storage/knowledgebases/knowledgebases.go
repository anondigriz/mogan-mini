package knowledgebases

import (
	"context"
	"path"

	"go.uber.org/zap"
)

const (
	KnowledgeBasesSubDir = "knowledgebases"
)

type Storage struct {
	lg                *zap.Logger
	KnowledgeBasesDir string
}

func New(lg *zap.Logger, workspaceDir string) *Storage {
	st := &Storage{
		lg:                lg,
		KnowledgeBasesDir: path.Join(workspaceDir, KnowledgeBasesSubDir),
	}

	return st
}

func (st Storage) Shutdown() {

}

func (st Storage) Ping(ctx context.Context) error {
	return nil
}
