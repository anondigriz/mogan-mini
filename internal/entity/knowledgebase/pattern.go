package knowledgebase

type Pattern struct {
	BaseInfo
	Type      TypePattern
	ExtraData ExtraDataPattern
}

type TypePattern int

const (
	Program    TypePattern = 0
	Constraint TypePattern = 1
	Formula    TypePattern = 2
	IfThenElse TypePattern = 3
)

type ExtraDataPattern struct {
	Description      string
	Language         Language
	Script           string
	InputParameters  []ParameterPattern
	OutputParameters []ParameterPattern
}

type ParameterPattern struct {
	ShortName string
	Type      TypeParameter
}

type Language int

const (
	JS     Language = 0
	Python Language = 1
)
