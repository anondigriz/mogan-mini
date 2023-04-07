package knowledgebase

type Parameter struct {
	BaseInfo
	GroupID   string
	Type      TypeParameter
	ExtraData ExtraDataParameter
}

type TypeParameter int

const (
	String     TypeParameter = 0
	Double     TypeParameter = 1
	Boolean    TypeParameter = 2
	BigInteger TypeParameter = 3
)

type ExtraDataParameter struct {
	Description  string
	DefaultValue string
}
