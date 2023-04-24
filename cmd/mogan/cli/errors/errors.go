package errors

const (
	KnowledgeBaseNotChosenErrMsg = "The knowledge base is not chosen. Please select the knowledge base using the `kb choose` command."
	ShortNameIsEmptyErrMsg       = "Short name must not be empty"
	IDIsEmptyErrMsg              = "ID must not be empty"

	GetAllKnowledgeBasesErrMsg   = "fail to get the knowledge bases' projects from the base project directory"
	ShowTUIKnowledgeBasesErrMsg  = "fail to show list of the knowledge bases' projects"
	RunTUIProgramErrMsg          = "fail when interacting with the console"
	ChooseTUIKnowledgeBaseErrMsg = "fail to get a choice of the knowledge base"
	InputTUINameErrMsg           = "fail when entering the name of the knowledge base name"
	EditTUIKnowledgeBaseErrMsg   = "fail to edit the knowledge base information"

	ChooseKnowledgeBaseErrMsg            = "fail when choosing a knowledge base from the base project directory"
	ReceivedResponseWasNotExpectedErrMsg = "received a response form that was not expected"
	KnowledgeBaseWasNotChosenErrMsg      = "knowledge base was not chosen"
	UpdateConfigErrMsg                   = "fail to update config"
	NameWasNotEnteredErrMsg              = "name was not entered"
	CreateKnowledgeBaseProjectErrMsg     = "fail to create the knowledge base project"
	GetKnowledgeBaseErrMsg               = "fail to get the knowledge base information"
	UpdateKnowledgeBaseErrMsg            = "fail to update the knowledge base information"
	KnowledgeBaseWasNotEditedErrMsg      = "knowledge base was not edited"

	ShowErrorPattern = "\n---\n🤦‍♀️ %v\n"
)
