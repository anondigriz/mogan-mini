package errors

const (
	PrepareConnectionFail          = "PingFail"
	GetKnowledgeBaseStorageFail    = "GetKnowledgeBaseStorageFail"
	UpdateKnowledgeBaseStorageFail = "UpdateKnowledgeBaseStorageFail"
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
