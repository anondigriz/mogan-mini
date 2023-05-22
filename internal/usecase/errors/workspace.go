package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

const (
	MakeWorkspaceDirFail   = "MakeWorkspaceDirFail"
	NotSupportedConfigType = "NotSupportedConfigType"
	CreateCfgFileFail      = "CreateCfgFileFail"
)

func NewMakeWorkspaceDirFailErr(err error, workspaceDir string) error {
	return UseCaseErr{
		Stat:    MakeWorkspaceDirFail,
		Message: fmt.Sprintf("%s. Workspace directory: '%s'", errMsgs.MakeWorkspaceDirFail, workspaceDir),
		Err:     err,
		Dt: map[string]string{
			"path": workspaceDir,
		},
	}
}

func NewNotSupportedConfigTypeErr(cfgType string) error {
	return UseCaseErr{
		Stat:    NotSupportedConfigType,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.NotSupportedConfigType, cfgType),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewCreateCfgFileFailErr(err error, filePath string) error {
	return UseCaseErr{
		Stat:    CreateCfgFileFail,
		Message: fmt.Sprintf("%s. Path: '%s'", errMsgs.CreateCfgFileFail, filePath),
		Err:     err,
		Dt: map[string]string{
			"path": filePath,
		},
	}
}
