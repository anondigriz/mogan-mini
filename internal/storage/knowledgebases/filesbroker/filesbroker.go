package filesbroker

import (
	"go.uber.org/zap"
)

type FilesBroker struct {
	lg            *zap.Logger
	currentDir    string
	fileExtension string
}

func New(lg *zap.Logger, currentDir, fileExtension string) *FilesBroker {
	pm := &FilesBroker{
		lg:            lg,
		currentDir:    currentDir,
		fileExtension: fileExtension,
	}
	return pm
}
