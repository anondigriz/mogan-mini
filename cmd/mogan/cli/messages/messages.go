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
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("you have chosen the knowledge base project with UUID '%s'", uuid))
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

func PrintReceivedNewEntityInfo() {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, "you have entered new information about the entity")
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

func PrintChooseGroup() {
	fmt.Printf(MsgPattern, "Choose group")
}

func PrintChosenGroup(uuid string) {
	fmt.Printf(MsgPatternKaomoji, OkKaomoji, fmt.Sprintf("you have chosen the group with UUID '%s'", uuid))
}
