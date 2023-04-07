package mappers

import (
	"github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
)

type PatternRow struct {
	BaseInfoForRow
	KnowledgeBaseUUID string
	Type              int
	ExtraData         ExtraDataPatternForRow
}

type ExtraDataPatternForRow struct {
	Description      string                   `json:"description"`
	Language         int                      `json:"language"`
	Script           string                   `json:"script"`
	InputParameters  []ParameterPatternForRow `json:"inputParameters"`
	OutputParameters []ParameterPatternForRow `json:"outputParameters"`
}

type ParameterPatternForRow struct {
	ShortName string `json:"shortName"`
	Type      int    `json:"type"`
}

func (pp *ParameterPatternForRow) Fill(base knowledgebase.ParameterPattern) {
	pp.ShortName = base.ShortName
	pp.Type = int(base.Type)
}

func (ex *ExtraDataPatternForRow) Fill(base knowledgebase.ExtraDataPattern) {
	ex.Description = base.Description
	ex.Language = int(base.Language)
	ex.Script = base.Script
	ex.InputParameters = []ParameterPatternForRow{}
	for _, v := range base.InputParameters {
		pp := ParameterPatternForRow{}
		pp.Fill(v)
		ex.InputParameters = append(ex.InputParameters, pp)
	}
	ex.OutputParameters = []ParameterPatternForRow{}
	for _, v := range base.OutputParameters {
		pp := ParameterPatternForRow{}
		pp.Fill(v)
		ex.OutputParameters = append(ex.OutputParameters, pp)
	}
}

func (pr *PatternRow) Fill(base knowledgebase.Pattern) {
	pr.BaseInfoForRow.Fill(base.BaseInfo)
	pr.Type = int(base.Type)
	pr.ExtraData.Fill(base.ExtraData)
}

func (pp *ParameterPatternForRow) Extract() knowledgebase.ParameterPattern {
	p := knowledgebase.ParameterPattern{}
	p.ShortName = pp.ShortName
	p.Type = knowledgebase.TypeParameter(pp.Type)
	return p
}

func (ex *ExtraDataPatternForRow) Extract() knowledgebase.ExtraDataPattern {
	e := knowledgebase.ExtraDataPattern{}
	e.Description = ex.Description
	e.Language = knowledgebase.Language(ex.Language)
	e.Script = ex.Script
	e.InputParameters = []knowledgebase.ParameterPattern{}
	for _, v := range ex.InputParameters {
		e.InputParameters = append(e.InputParameters, v.Extract())
	}
	e.OutputParameters = []knowledgebase.ParameterPattern{}
	for _, v := range ex.OutputParameters {
		e.OutputParameters = append(e.OutputParameters, v.Extract())
	}
	return e
}

func (pr *PatternRow) Extract() knowledgebase.Pattern {
	k := knowledgebase.Pattern{}
	k.BaseInfo = pr.BaseInfoForRow.Extract()
	k.Type = knowledgebase.TypePattern(pr.Type)
	k.ExtraData = pr.ExtraData.Extract()
	return k
}
