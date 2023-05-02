package errors

import "fmt"

func NewXMLUnmarshalFailErr(e error) error {
	return UseCaseErr{
		Stat:    XMLUnmarshalFail,
		Message: "fail to unmarshal the xml file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewParsingXMLFailErr(msg string, e error) error {
	return UseCaseErr{
		Stat:    ParsingXMLFail,
		Message: msg,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewReadingXMLFailErr(e error) error {
	return UseCaseErr{
		Stat:    ReadingXMLFail,
		Message: "fail to read the XML file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewUnsupportedFormatXMLVersionErr(version string) error {
	return UseCaseErr{
		Stat:    UnsupportedFormatXMLVersion,
		Message: fmt.Sprintf("xml exchange document file version '%s' is not supported", version),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnexpectedJobExecutionFailErr(err error) error {
	return UseCaseErr{
		Stat:    UnexpectedJobExecutionFail,
		Message: "unexpected job execution error occurred",
		Err:     nil,
		Dt:      map[string]string{},
	}
}
