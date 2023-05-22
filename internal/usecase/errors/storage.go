package errors

const (
	StorageFail = "StorageFail"
)

type storageErr interface {
	error
	Status() string
	Data() map[string]string
	IsStorageErr() bool
}

func WrapStorageFailErr(err error) error {
	switch e := err.(type) {
	case storageErr:
		{
			switch e.Status() {
			default:
				return NewStorageFailErr(e)
			}
		}
	}
	return NewUnexpectedUseCaseFailErr(err)
}

func NewStorageFailErr(err storageErr) error {
	return UseCaseErr{
		Stat:    StorageFail,
		Message: err.Error(),
		Err:     err,
		Dt:      err.Data(),
	}
}
