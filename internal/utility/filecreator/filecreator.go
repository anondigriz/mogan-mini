package filecreator

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

type FileCreator struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *FileCreator {
	fc := &FileCreator{
		lg: lg,
	}
	return fc
}

func (fc *FileCreator) CreateFile(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			fc.lg.Error(fmt.Sprintf("fail to create a file on the path %s", filePath),
				zap.Error(err), zap.Reflect("path", filePath))
			return err
		}
		defer file.Close()
	}
	return nil
}
