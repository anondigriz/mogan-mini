package filesbroker

import (
	"os"
)

func (fb FilesBroker) IsFileExistByUUID(uuid string) bool {
	filePath := fb.GetFilePath(uuid)
	return fb.IsFileExistByPath(filePath)
}

func (fb FilesBroker) IsFileExistByPath(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}
