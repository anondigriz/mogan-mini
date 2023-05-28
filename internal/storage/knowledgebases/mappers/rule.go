package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

type Rule struct {
	BaseInfo
	PatternUUID      string
	InputParameters  []ParameterRule
	OutputParameters []ParameterRule
}

type ParameterRule struct {
	ShortName     string
	ParameterUUID string
}

func (r *Rule) Fill(base kbEnt.Rule) {
	r.BaseInfo.Fill(base.BaseInfo)
	r.PatternUUID = base.PatternUUID

	r.InputParameters = make([]ParameterRule, len(base.InputParameters))
	for _, v := range base.InputParameters {
		var pp ParameterRule
		pp.Fill(v)
		r.InputParameters = append(r.InputParameters, pp)
	}

	r.OutputParameters = make([]ParameterRule, len(base.OutputParameters))
	for _, v := range base.OutputParameters {
		var pp ParameterRule
		pp.Fill(v)
		r.OutputParameters = append(r.OutputParameters, pp)
	}
}

func (p Rule) Extract() kbEnt.Rule {
	result := kbEnt.Rule{
		BaseInfo: p.BaseInfo.Extract(),
	}
	result.PatternUUID = p.PatternUUID

	result.InputParameters = make([]kbEnt.ParameterRule, len(p.InputParameters))
	for _, v := range p.InputParameters {
		result.InputParameters = append(result.InputParameters, v.Extract())
	}

	result.OutputParameters = make([]kbEnt.ParameterRule, len(p.OutputParameters))
	for _, v := range p.OutputParameters {
		result.OutputParameters = append(result.OutputParameters, v.Extract())
	}

	return result
}

func (p *ParameterRule) Fill(base kbEnt.ParameterRule) {
	p.ShortName = base.ShortName
	p.ParameterUUID = base.ParameterUUID
}

func (p ParameterRule) Extract() kbEnt.ParameterRule {
	result := kbEnt.ParameterRule{
		ShortName:     p.ShortName,
		ParameterUUID: p.ParameterUUID,
	}
	return result
}
