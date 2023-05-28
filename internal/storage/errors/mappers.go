package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

const (
	TypeIsNotSupportedByStorage = "TypeIsNotSupportedByStorage"
)

func NewTypeIsNotSupportedByStorageErr(t string) error {
	return StorageErr{
		Stat:    TypeIsNotSupportedByStorage,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.TypeIsNotSupportedByStorage, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}
