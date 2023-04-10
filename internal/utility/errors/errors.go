package errors

const (
	ReadingXMLFail                   = "ReadingXMLFail"
	UnsupportedFormatXMLVersion      = "UnsupportedFormatXMLVersion"
	UnknownXMLValidationError        = "UnknownXMLValidationError"
	InsertKnowledgeBaseToStorageFail = "InsertKnowledgeBaseToStorageFail"
	InsertXMLFileToStorageFail       = "InsertXMLFileToStorageFail"
	InsertXMLParseJobToStorageFail   = "InsertXMLParseJobToStorageFail"
	UnexpectedStorageFail            = "UnexpectedStorageFail"
	NoDataFound                      = "NoDataFound"
	NotPartOfKnowledgeBase           = "NotPartOfKnowledgeBase"
	XMLUnmarshalFail                 = "XMLUnmarshalFail"
	ParsingXMLFail                   = "ParsingXMLFail"
	UnexpectedJobExecutionFail       = "UnexpectedJobExecutionFail"
)

type UtilityErr struct {
	Stat    string
	Message string
	Err     error
	Dt      map[string]string
}

func (er UtilityErr) Status() string {
	return er.Stat
}

func (er UtilityErr) Error() string {
	return er.Message
}

func (er UtilityErr) Data() map[string]string {
	return er.Dt
}

func (er UtilityErr) Unwrap() error {
	return er.Err
}

func (er UtilityErr) FromUtility() bool {
	return true
}
