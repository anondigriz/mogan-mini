package messages

import "fmt"

const (
	OkKaomoji       = "d-(´▽｀)-b"
	FailKaomoji     = "¯\\_(⊙︿⊙)_/¯"
	DontKnowKaomoji = "¯\\_(ツ)_/¯"

	MsgPattern        = "\n---\n%v\n"
	MsgPatternKaomoji = "\n---\n%v %v\n"
)

func PrintChosenKnowledgeBase(uuid string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("you have chosen the knowledge base with UUID '%s'", uuid))
}

func PrintEnteredShortNameKnowledgeBase(shortName string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("you have entered the short name of the knowledge base: '%s'", shortName))
}

func PrintEnteredIDKnowledgeBase(id string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("you have entered the ID of the knowledge base:'%s'", id))
}

func PrintCreatedKnowledgeBase(uuid string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("knowledge base project has been created with UUID '%s'", uuid))
}

func PrintReceivedNewObjectInfo() {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, "you have entered new information about the object")
}

func PrintKnowledgeBaseRemoved(uuid string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("knowledge base project with UUID '%s' successfully deleted", uuid))
}

func PrintNoDataToShow() {
	fmt.Printf(MsgPatternKaomoji, DontKnowKaomoji, "no data to display")
}

func PrintChangesAccepted() {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, "changes accepted")
}

func PrintFail(msg string) {
	fmt.Printf(MsgPatternKaomoji, FailKaomoji, msg)
}

func PrintChosenGroup(uuid string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("you have chosen the group with UUID '%s'", uuid))
}

func PrintBaseInfoNotEdited() {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, "base info about the object was not edited")
}

func PrintChooseGroup() {
	fmt.Printf(MsgPattern, "Choose group")
}

func PrintChooseKnowledgeBase() {
	fmt.Printf(MsgPattern, "Choose knowledge base")
}
func PrintEditDescription() {
	fmt.Printf(MsgPattern, "Edit the description of the object")
}

func PrintEditIDShortName() {
	fmt.Printf(MsgPattern, "Edit the ID and short name of the object")
}

func PrintConfirmRemoveKnowledgeBase() {
	fmt.Printf(MsgPattern, "Confirm the removing of the local knowledge base project. This action cannot be undone.")
}
