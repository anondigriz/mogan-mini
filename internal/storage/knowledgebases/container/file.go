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
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return []byte{}, errors.NewReadFileFailErr(err, fb.GetFilePath(uuid))
	}
	return data, nil
}

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

func (c Container) removeFile(uuid string, subDir string) error {
	fb := filesbroker.New(c.lg, subDir, fileExtension)
	err := fb.RemoveFileByUUID(uuid)
	if err != nil {
		c.lg.Error(errMsgs.DeleteFileFail, zap.Error(err))
		return err
	}
	return nil
}

func (c Container) getFilesUUIDsInDir(subDir string) []string {
	dir := path.Join(c.knowledgeBaseDir, subDir)
	fb := filesbroker.New(c.lg, dir, fileExtension)
	paths := fb.GetAllFilesPaths()
	result := make([]string, 0, len(paths))
	for _, v := range paths {
		uuid := fb.GetFileUUID(v)
		result = append(result, uuid)
	}
	return result
}
