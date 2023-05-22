package filesbroker

import (
	"os"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (fb FilesBroker) CreateFileByUUID(uuid string) (*os.File, error) {
	filePath := fb.GetFilePath(uuid)
	return fb.CreateFileByPath(filePath)
}

func (fb FilesBroker) CreateFileByPath(filePath string) (*os.File, error) {
	err := os.MkdirAll(fb.settings.KnowledgeBaseDir, os.ModePerm)
	if err != nil {
		fb.lg.Error(errMsgs.CreateKnowledgeBaseFileFail, zap.Error(err))
		return nil, errors.NewCreateKnowledgeBaseFileFailErr(err, filePath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		fb.lg.Error(errMsgs.CreateKnowledgeBaseFileFail, zap.Error(err))
		return nil, errors.NewCreateKnowledgeBaseFileFailErr(err, filePath)
	}
	return file, nil
}
