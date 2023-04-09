package errors

import "fmt"

var (
	KnowledgeBaseNotChosenErr = fmt.Errorf("The knowledge base is not chosen. Please select the knowledge base using the `kb choose` command.")
	ShortNameIsEmptyErr       = fmt.Errorf("Short name must not be empty")
	IDIsEmptyErr              = fmt.Errorf("ID must not be empty")
)
