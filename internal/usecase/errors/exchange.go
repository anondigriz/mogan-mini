package errors

import (
	"errors"

	exchangeErr "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/errors"
)

func WrapXMLValidationErr(err error) error {
	var e exchangeErr.ValidatorErr
	if errors.As(err, &e) {
		return UseCaseErr{
			Stat:    e.Stat,
			Message: e.Message,
			Err:     e,
			Dt:      e.Dt,
		}
	}
	return UseCaseErr{
		Stat:    UnknownXMLValidationError,
		Message: "unknown XML validation error has occurred",
		Err:     err,
		Dt:      map[string]string{},
	}
}

func NewNotPartOfKnowledgeBaseErr() error {
	return UseCaseErr{
		Stat:    NotPartOfKnowledgeBase,
		Message: "entity is not part of the knowledge base",
		Err:     nil,
		Dt:      map[string]string{},
	}
}
