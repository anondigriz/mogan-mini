package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

const (
	KnowledgeBaseFileNotFound   = "KnowledgeBaseFileNotFound"
	WalkInWorkspaceDirFail      = "WalkInWorkspaceDirFail"
	DeleteKnowledgeBaseFileFail = "DeleteKnowledgeBaseFileFail"
	CreateKnowledgeBaseFileFail = "CreateKnowledgeBaseFileFail"
	OpenKnowledgeBaseFileFail   = "OpenKnowledgeBaseFileFail"
)

func NewKnowledgeBaseFileNotFoundErr(err error, filePath string) error {
	return StorageErr{
		Stat:    KnowledgeBaseFileNotFound,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.KnowledgeBaseFileNotFound, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewWalkInWorkspaceDirFailErr(err error, workspaceDir string) error {
	return StorageErr{
		Stat:    WalkInWorkspaceDirFail,
		Message: fmt.Sprintf("%s. Workspace directory: '%s'", errMsgs.WalkInWorkspaceDirFail, workspaceDir),
		Err:     err,
		Dt: map[string]string{
			"path": workspaceDir,
		},
	}
}

func NewDeleteKnowledgeBaseFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    DeleteKnowledgeBaseFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.DeleteKnowledgeBaseFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewCreateKnowledgeBaseFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    CreateKnowledgeBaseFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.CreateKnowledgeBaseFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}

func NewOpenKnowledgeBaseFileFailErr(err error, filePath string) error {
	return StorageErr{
		Stat:    OpenKnowledgeBaseFileFail,
		Message: fmt.Sprintf("%s. File path: '%s'", errMsgs.OpenKnowledgeBaseFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}
