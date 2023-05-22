package filesbroker

import (
	"path"

	"go.uber.org/zap"
)

const (
	fileExtension = ".xml"
	subDir        = "knowledgebases"
)

type settings struct {
	KnowledgeBaseDir string
}

type FilesBroker struct {
	lg       *zap.Logger
	settings settings
}

func New(lg *zap.Logger, workspaceDir string) *FilesBroker {
	pm := &FilesBroker{
		settings: settings{
			KnowledgeBaseDir: path.Join(workspaceDir, subDir),
		},
		lg: lg,
	}
	return pm
}
