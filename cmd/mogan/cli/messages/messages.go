package messages

import "fmt"

const (
	OkKaomoji       = "d-(´▽｀)-b"
	FailKaomoji     = "¯\\_(⊙︿⊙)_/¯"
	DontKnowKaomoji = "¯\\_(ツ)_/¯"

	MsgPattern = "\n---\n%v %v\n"
)

func PrintChosenKnowledgeBase(uuid string) {
	fmt.Printf(MsgPattern, OkKaomoji, fmt.Sprintf("you have chosen the knowledge base project with UUID '%s'", uuid))
}

func PrintEnteredShortNameKnowledgeBase(shortName string) {
	fmt.Printf(MsgPattern, OkKaomoji, fmt.Sprintf("you have entered the short name of the knowledge base: '%s'", shortName))
}

func PrintEnteredIDKnowledgeBase(id string) {
	fmt.Printf(MsgPattern, OkKaomoji, fmt.Sprintf("you have entered the ID of the knowledge base:'%s'", id))
}

func PrintCreatedKnowledgeBase(uuid string) {
	fmt.Printf(MsgPattern, OkKaomoji, fmt.Sprintf("knowledge base project has been created with UUID '%s'", uuid))
}

func PrintReceivedNewEntityInfo() {
	fmt.Printf(MsgPattern, OkKaomoji, "you have entered new information about the entity")
}

func PrintKnowledgeBaseRemoved(uuid string) {
	fmt.Printf(MsgPattern, OkKaomoji, fmt.Sprintf("knowledge base project with UUID '%s' successfully deleted", uuid))
}

func PrintNoDataToShow() {
	fmt.Printf(MsgPattern, DontKnowKaomoji, "no data to display")

}

func PrintChangesAccepted() {
	fmt.Printf(MsgPattern, OkKaomoji, "changes accepted")
}

func PrintFail(msg string) {
	fmt.Printf(MsgPattern, FailKaomoji, msg)
}
