package errors

const (
	PingFail              = "PingFail"
	ReadFileFail          = "ReadFileFail"
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

func (se StorageErr) FromStorage() bool {
	return true
}

func NewPingFailErr(err error) error {
	return StorageErr{
		Stat:    PingFail,
		Message: "storage ping fail",
		Err:     err,
		Dt:      map[string]string{},
	}
}
