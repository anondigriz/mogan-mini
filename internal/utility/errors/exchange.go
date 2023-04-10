package errors

import (
	"errors"

	validatorErr "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/validator/errors"
)

func WrapXMLValidationErr(err error) error {
	var e validatorErr.ValidatorErr
	if errors.As(err, &e) {
		return UtilityErr{
			Stat:    e.Stat,
			Message: e.Message,
			Err:     e,
			Dt:      e.Dt,
		}
	}
	return UtilityErr{
		Stat:    UnknownXMLValidationError,
		Message: "unknown XML validation error has occurred",
		Err:     err,
		Dt:      map[string]string{},
	}
}

func NewNotPartOfKnowledgeBaseErr() error {
	return UtilityErr{
		Stat:    NotPartOfKnowledgeBase,
		Message: "entity is not part of the knowledge base",
		Err:     nil,
		Dt:      map[string]string{},
	}
}
