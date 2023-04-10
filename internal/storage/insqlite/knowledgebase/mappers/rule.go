package mappers

import (
	"encoding/json"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type RuleRow struct {
	BaseInfoForRow
	PatternUUID string
	ExtraData   string
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

func (pr *ParameterRuleForRow) Fill(base kbEnt.ParameterRule) {
	pr.ShortName = base.ShortName
	pr.ParameterUUID = base.ParameterUUID
}

func (ex *ExtraDataRuleForRow) Fill(base kbEnt.ExtraDataRule) {
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

func (rr *RuleRow) Fill(base kbEnt.Rule) error {
	rr.BaseInfoForRow.Fill(base.BaseInfo)
	rr.PatternUUID = base.PatternUUID
	err := rr.fillExtraData(base)
	if err != nil {
		return err
	}
	return nil
}

func (rr *RuleRow) fillExtraData(base kbEnt.Rule) error {
	ex := ExtraDataRuleForRow{}
	ex.Fill(base.ExtraData)

	b, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	rr.ExtraData = string(b)
	return nil
}

func (pr ParameterRuleForRow) Extract() kbEnt.ParameterRule {
	p := kbEnt.ParameterRule{}
	p.ShortName = pr.ShortName
	p.ParameterUUID = pr.ParameterUUID
	return p
}

func (ex ExtraDataRuleForRow) Extract() kbEnt.ExtraDataRule {
	e := kbEnt.ExtraDataRule{}
	e.Description = ex.Description
	e.InputParameters = []kbEnt.ParameterRule{}
	for _, v := range ex.InputParameters {
		e.InputParameters = append(e.InputParameters, v.Extract())
	}
	e.OutputParameters = []kbEnt.ParameterRule{}
	for _, v := range ex.OutputParameters {
		e.OutputParameters = append(e.OutputParameters, v.Extract())
	}
	return e
}

func (rr RuleRow) Extract() (kbEnt.Rule, error) {
	r := kbEnt.Rule{}
	r.BaseInfo = rr.BaseInfoForRow.Extract()
	r.PatternUUID = rr.PatternUUID

	ex, err := rr.extractExtraData()
	if err != nil {
		return kbEnt.Rule{}, err
	}
	r.ExtraData = ex
	return r, nil
}

func (rr RuleRow) extractExtraData() (kbEnt.ExtraDataRule, error) {
	ex := &ExtraDataRuleForRow{}
	err := json.Unmarshal([]byte(rr.ExtraData), &ex)
	if err != nil {
		return kbEnt.ExtraDataRule{}, nil
	}
	return ex.Extract(), nil
}
