package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

const (
	ObjectNotFound = "ObjectNotFound"
)

func NewObjectNotFoundErr(uuid string) error {
	return UseCaseErr{
		Stat:    ObjectNotFound,
		Message: fmt.Sprintf("%s. UUID: '%s'", errMsgs.ObjectNotFound, uuid),
		Err:     nil,
		Dt: map[string]string{
			"uuid": uuid,
		},
	}
}
