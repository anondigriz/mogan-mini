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
	err := fb.CreateDir(fb.currentDir)
	if err != nil {
		fb.lg.Error(errMsgs.CreateFileFail, zap.Error(err))
		return nil, errors.NewCreateFileFailErr(err, filePath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		fb.lg.Error(errMsgs.CreateFileFail, zap.Error(err))
		return nil, errors.NewCreateFileFailErr(err, filePath)
	}
	return file, nil
}

func (fb FilesBroker) CreateDir(dirPath string) error {
	err := os.MkdirAll(fb.currentDir, os.ModePerm)
	if err != nil {
		fb.lg.Error(errMsgs.CreateDirFail, zap.Error(err))
		return errors.NewCreateDirFailErr(err, dirPath)
	}
	return nil
}
