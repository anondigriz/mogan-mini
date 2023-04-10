package mappers

import (
	"github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type RuleRow struct {
	BaseInfoForRow
	PatternUUID string
	ExtraData   ExtraDataRuleForRow
}

type ExtraDataRuleForRow struct {
	Description      string                `json:"description"`
	InputParameters  []ParameterRuleForRow `json:"inputParameters"`
	OutputParameters []ParameterRuleForRow `json:"outputParameters"`
}

type ParameterRuleForRow struct {
	ShortName     string `json:"shortName"`
	ParameterUUID string `json:"parameterID"`
}

func (pr *ParameterRuleForRow) Fill(base knowledgebase.ParameterRule) {
	pr.ShortName = base.ShortName
	pr.ParameterUUID = base.ParameterUUID
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
	rr.PatternUUID = base.PatternUUID
	rr.ExtraData.Fill(base.ExtraData)
}

func (pr *ParameterRuleForRow) Extract() knowledgebase.ParameterRule {
	p := knowledgebase.ParameterRule{}
	p.ShortName = pr.ShortName
	p.ParameterUUID = pr.ParameterUUID
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
	k.PatternUUID = rr.PatternUUID
	k.ExtraData = rr.ExtraData.Extract()
	return k
}
