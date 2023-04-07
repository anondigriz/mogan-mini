package mappers

import (
	"github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
)

type RuleRow struct {
	BaseInfoForRow
	PatternID string
	ExtraData ExtraDataRuleForRow
}

type ExtraDataRuleForRow struct {
	Description      string                `json:"description"`
	InputParameters  []ParameterRuleForRow `json:"inputParameters"`
	OutputParameters []ParameterRuleForRow `json:"outputParameters"`
}

type ParameterRuleForRow struct {
	ShortName   string `json:"shortName"`
	ParameterID string `json:"parameterID"`
}

func (pr *ParameterRuleForRow) Fill(base knowledgebase.ParameterRule) {
	pr.ShortName = base.ShortName
	pr.ParameterID = base.ParameterID
}

func (ex *ExtraDataRuleForRow) Fill(base knowledgebase.ExtraDataRule) {
	ex.Description = base.Description
	ex.InputParameters = []ParameterRuleForRow{}
	for _, v := range base.InputParameters {
		pp := ParameterRuleForRow{}
		pp.Fill(v)
		ex.InputParameters = append(ex.InputParameters, pp)
	}
	ex.OutputParameters = []ParameterRuleForRow{}
	for _, v := range base.OutputParameters {
		pp := ParameterRuleForRow{}
		pp.Fill(v)
		ex.OutputParameters = append(ex.OutputParameters, pp)
	}
}

func (rr *RuleRow) Fill(base knowledgebase.Rule) {
	rr.BaseInfoForRow.Fill(base.BaseInfo)
	rr.PatternID = base.PatternID
	rr.ExtraData.Fill(base.ExtraData)
}

func (pr *ParameterRuleForRow) Extract() knowledgebase.ParameterRule {
	p := knowledgebase.ParameterRule{}
	p.ShortName = pr.ShortName
	p.ParameterID = pr.ParameterID
	return p
}

func (ex *ExtraDataRuleForRow) Extract() knowledgebase.ExtraDataRule {
	e := knowledgebase.ExtraDataRule{}
	e.Description = ex.Description
	e.InputParameters = []knowledgebase.ParameterRule{}
	for _, v := range ex.InputParameters {
		e.InputParameters = append(e.InputParameters, v.Extract())
	}
	e.OutputParameters = []knowledgebase.ParameterRule{}
	for _, v := range ex.OutputParameters {
		e.OutputParameters = append(e.OutputParameters, v.Extract())
	}
	return e
}
func (rr *RuleRow) Extract() knowledgebase.Rule {
	k := knowledgebase.Rule{}
	k.BaseInfo = rr.BaseInfoForRow.Extract()
	k.PatternID = rr.PatternID
	k.ExtraData = rr.ExtraData.Extract()
	return k
}
