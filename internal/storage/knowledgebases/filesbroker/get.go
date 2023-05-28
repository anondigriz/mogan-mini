package filesbroker

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (fb FilesBroker) GetFilePath(uuid string) string {
	return path.Join(fb.currentDir, uuid+fb.fileExtension)
}

func (fb FilesBroker) GetFileUUID(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

func (fb FilesBroker) GetAllFilesPaths() []string {
	var paths []string
	filepath.WalkDir(fb.currentDir, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			fb.lg.Error(errMsgs.WalkInDirFail, zap.Error(e))
			return errors.NewWalkInDirFailErr(e, fb.currentDir)
		}
		if filepath.Ext(d.Name()) == fb.fileExtension {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}

func (fb FilesBroker) GetAllChildDirNames() []string {
	var paths []string
	entries, err := os.ReadDir(fb.currentDir)
	if err != nil {
		fb.lg.Error(errMsgs.WalkInDirFail, zap.Error(err))
		return []string{}
	}
	for _, e := range entries {
		if !e.IsDir() {
			paths = append(paths, e.Name())
		}
	}
	return paths
}
