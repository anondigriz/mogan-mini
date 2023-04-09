package knowledgebase

type KnowledgeBase struct {
	BaseInfo
	RemoteUUID string
	ExtraData  ExtraDataKnowledgeBase
	Path       string
}

type ExtraDataKnowledgeBase struct {
	Description string
	Groups      GroupHierarchy
}

type GroupHierarchy struct {
	GroupUUID string
	Contains  []GroupHierarchy
}
