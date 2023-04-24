package errors

func NewPrepareConnectionErr(err error) error {
	return UseCaseErr{
		Stat:    PrepareConnectionFail,
		Message: "fail to prepare connection to the knowledge base project's storage",
		Err:     err,
		Dt:      map[string]string{},
	}
}
