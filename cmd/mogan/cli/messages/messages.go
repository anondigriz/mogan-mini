package messages

import "fmt"

const (
	OkEmoji       = "👍"
	FailEmoji     = "🤦‍"
	DontKnowEmoji = "🤷"
	
	MsgPattern = "\n---\n%v %v\n"
)

func PrintChosenKnowledgeBase(uuid string) {
	fmt.Printf(MsgPattern, OkEmoji, fmt.Sprintf("you have chosen the knowledge base project with UUID '%s'", uuid))
}

func PrintEnteredShortNameKnowledgeBase(shortName string) {
	fmt.Printf(MsgPattern, OkEmoji, fmt.Sprintf("you have entered the knowledge base name '%s'", shortName))
}

func PrintCreatedKnowledgeBase(uuid string) {
	fmt.Printf(MsgPattern, OkEmoji, fmt.Sprintf("knowledge base project has been created with UUID '%s'", uuid))
}

func PrintRecivedNewEntityInfo() {
	fmt.Printf(MsgPattern, OkEmoji, "you have entered new information about the entity")
}

func PrintKnowledgeBaseRemoved(uuid string) {
	fmt.Printf(MsgPattern, OkEmoji, fmt.Sprintf("knowledge base project with UUID '%s' successfully deleted", uuid))

}

func PrintNoDataToShow() {
	fmt.Printf(MsgPattern, DontKnowEmoji, "no data to display")

}

func PrintChangesAccepted() {
	fmt.Printf(MsgPattern, OkEmoji, "changes accepted")
}

func PrintFail(msg string) {
	fmt.Printf(MsgPattern, FailEmoji, msg)
}
