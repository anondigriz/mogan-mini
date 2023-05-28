package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

const (
	FileNotFound   = "FileNotFound"
	WalkInDirFail  = "WalkInDirFail"
	DeleteFileFail = "DeleteFileFail"
	CreateFileFail = "CreateFileFail"
	OpenFileFail   = "OpenFileFail"
	CreateDirFail  = "CreateDirFail"
	WriteFileFail  = "WriteFileFail"
	ReadFileFail   = "ReadFileFail"
)

func NewFileNotFoundErr(err error, filePath string) error {
	return StorageErr{
		Stat:    FileNotFound,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.FileNotFound, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewWalkInDirFailErr(err error, dirPath string) error {
	return StorageErr{
		Stat:    WalkInDirFail,
		Message: fmt.Sprintf("%s. Directory: '%s'", errMsgs.WalkInDirFail, dirPath),
		Err:     err,
		Dt: map[string]string{
			"path": dirPath,
		},
	}
}

func NewDeleteFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    DeleteFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.DeleteFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewCreateFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    CreateFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.CreateFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewCreateDirFailErr(err error, dirPath string) error {
	return StorageErr{
		Stat:    CreateDirFail,
		Message: fmt.Sprintf("%s. Directory: '%s'", errMsgs.CreateDirFail, dirPath),
		Err:     err,
		Dt: map[string]string{
			"path": dirPath,
		},
	}
}

func NewOpenFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    OpenFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.OpenFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewWriteFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    WriteFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.WriteFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewReadFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    ReadFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.ReadFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}
