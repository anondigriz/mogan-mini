package filesbroker

import (
	"os"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (fb FilesBroker) OpenFileByUUID(uuid string) (*os.File, error) {
	filePath := fb.GetFilePath(uuid)
	return fb.OpenFileByPath(filePath)
}

func (fb FilesBroker) OpenFileByPath(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fb.lg.Error(errMsgs.OpenKnowledgeBaseFileFail, zap.Error(err))
		return nil, errors.NewOpenKnowledgeBaseFileFailErr(err, filePath)
	}

	return file, nil
}
