package errors

func NewGetKnowledgeBaseStorageErr(err error) error {
	return UseCaseErr{
		Stat:    GetKnowledgeBaseStorageFail,
		Message: "fail to get the knowledge base from the storage",
		Err:     err,
		Dt:      map[string]string{},
	}
}

func NewUpdateKnowledgeBaseStorageErr(err error) error {
	return UseCaseErr{
		Stat:    UpdateKnowledgeBaseStorageFail,
		Message: "fail to update the knowledge base in the storage",
		Err:     err,
		Dt:      map[string]string{},
	}
}
