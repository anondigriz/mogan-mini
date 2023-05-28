package filesbroker

import (
	"os"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (fb FilesBroker) RemoveFileByUUID(uuid string) error {
	filePath := fb.GetFilePath(uuid)
	return fb.RemoveFileByPath(filePath)
}

func (fb FilesBroker) RemoveFileByPath(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		fb.lg.Error(errMsgs.DeleteFileFail, zap.Error(err))
		return errors.NewDeleteFileFailErr(err, filePath)
	}
	return nil
}
