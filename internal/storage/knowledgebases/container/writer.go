package container

import (
	"path"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

func (c Container) writeFile(data []byte, uuid, subDir string) error {
	dir := path.Join(c.knowledgeBaseDir, subDir)
	fb := filesbroker.New(c.lg, dir, fileExtension)

	file, err := fb.CreateFileByUUID(uuid)
	if err != nil {
		c.lg.Error(errMsgs.CreateFileFail, zap.Error(err))
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return errors.NewWriteFileFailErr(err, fb.GetFilePath(uuid))
	}
	return nil
}
