package knowledgebase

type Rule struct {
	BaseInfo
	PatternID string
	ExtraData ExtraDataRule
}

type ExtraDataRule struct {
	Description      string
	InputParameters  []ParameterRule
	OutputParameters []ParameterRule
}

type ParameterRule struct {
	ShortName     string
	ParameterUUID string
}
