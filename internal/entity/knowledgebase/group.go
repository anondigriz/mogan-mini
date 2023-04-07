package knowledgebase

type Group struct {
	BaseInfo
	ExtraData ExtraDataGroup
}

type ExtraDataGroup struct {
	Description string
}
