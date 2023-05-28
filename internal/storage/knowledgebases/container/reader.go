package container

import (
	"io"
	"path"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

func (c Container) readFile(uuid, subDir string) ([]byte, error) {
	dir := path.Join(c.knowledgeBaseDir, subDir)
	fb := filesbroker.New(c.lg, dir, fileExtension)

	file, err := fb.OpenFileByUUID(uuid)
	if err != nil {
		c.lg.Error(errMsgs.CreateFileFail, zap.Error(err))
		return []byte{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		c.lg.Error(errMsgs.ReadFail, zap.Error(err))
		return []byte{}, errors.NewReadFailErr(err, fb.GetFilePath(uuid))
	}
	return data, nil
}
