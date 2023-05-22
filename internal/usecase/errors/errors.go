package errors

import (
	errMsgs "github.com/anondigriz/mogan-mini/internal/usecase/errors/messages"
)

const (
	UnexpectedUseCaseFail = "UnexpectedUseCaseFail"
)

type UseCaseErr struct {
	Stat    string
	Message string
	Err     error
	Dt      map[string]string
}

func (se UseCaseErr) Status() string {
	return se.Stat
}

func (se UseCaseErr) Error() string {
	return se.Message
}

func (se UseCaseErr) Data() map[string]string {
	return se.Dt
}

func (se UseCaseErr) Unwrap() error {
	return se.Err
}

func (se UseCaseErr) FromUseCase() bool {
	return true
}

func NewUnexpectedUseCaseFailErr(err error) error {
	return UseCaseErr{
		Stat:    ExchangeKnowledgeBaseFail,
		Message: errMsgs.UnexpectedUseCaseFail,
		Err:     err,
		Dt:      map[string]string{},
	}
}
