package errors

import (
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

const (
	PingFail              = "PingFail"
	NoDataFound           = "NoDataFound"
	UnexpectedStorageFail = "UnexpectedStorageFail"
)

type StorageErr struct {
	Stat    string
	Message string
	Err     error
	Dt      map[string]string
}

func (se StorageErr) Status() string {
	return se.Stat
}

func (se StorageErr) Error() string {
	return se.Message
}

func (se StorageErr) Data() map[string]string {
	return se.Dt
}

func (se StorageErr) Unwrap() error {
	return se.Err
}

func (se StorageErr) IsStorageErr() bool {
	return true
}

func NewPingFailErr(err error) error {
	return StorageErr{
		Stat:    PingFail,
		Message: errMsgs.PingFail,
		Err:     err,
		Dt:      map[string]string{},
	}
}

func NewNoDataFoundErr(err error, msg string) error {
	return StorageErr{
		Stat:    NoDataFound,
		Message: msg,
		Err:     err,
		Dt:      map[string]string{},
	}
}

func NewUnexpectedStorageFailErr(err error) error {
	return StorageErr{
		Stat:    UnexpectedStorageFail,
		Message: errMsgs.UnexpectedStorageFail,
		Err:     err,
		Dt:      map[string]string{},
	}
}
