package errors

import (
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

const (
	TomlMarshalFail   = "TomlMarshalFail"
	TomlUnmarshalFail = "TomlUnmarshalFail"
)

func NewTomlMarshalFailErr(err error) error {
	return StorageErr{
		Stat:    TomlMarshalFail,
		Message: errMsgs.TomlMarshalFail,
		Err:     err,
		Dt:      map[string]string{},
	}
}

func NewTomlUnmarshalFailErr(err error) error {
	return StorageErr{
		Stat:    TomlUnmarshalFail,
		Message: errMsgs.TomlUnmarshalFail,
		Err:     err,
		Dt:      map[string]string{},
	}
}
