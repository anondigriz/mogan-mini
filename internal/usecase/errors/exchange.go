package errors

const (
	ExchangeKnowledgeBaseFail = "ExchangeKnowledgeBaseFail"
)

type exchangeKnowledgeBaseErr interface {
	error
	Status() string
	Data() map[string]string
	IsExchangeKnowledgeBaseErr() bool
}

func WrapExchangeErr(err error) error {
	switch e := err.(type) {
	case exchangeKnowledgeBaseErr:
		{
			switch e.Status() {
			default:
				return NewExchangeKnowledgeBaseFailErr(e)
			}
		}
	}
	return NewUnexpectedUseCaseFailErr(err)
}

func NewExchangeKnowledgeBaseFailErr(err exchangeKnowledgeBaseErr) error {
	return UseCaseErr{
		Stat:    ExchangeKnowledgeBaseFail,
		Message: err.Error(),
		Err:     err,
		Dt:      err.Data(),
	}
}
