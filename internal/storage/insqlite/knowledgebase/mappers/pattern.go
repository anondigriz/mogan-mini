package mappers

import (
	"encoding/json"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type PatternRow struct {
	BaseInfoForRow
	Type      int
	ExtraData string
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

func (pp *ParameterPatternForRow) Fill(base kbEnt.ParameterPattern) {
	pp.ShortName = base.ShortName
	pp.Type = int(base.Type)
}

func (ex *ExtraDataPatternForRow) Fill(base kbEnt.ExtraDataPattern) {
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

func (pr *PatternRow) Fill(base kbEnt.Pattern) error {
	pr.BaseInfoForRow.Fill(base.BaseInfo)
	pr.Type = int(base.Type)
	err := pr.fillExtraData(base)
	if err != nil {
		return err
	}
	return nil
}

func (kr *PatternRow) fillExtraData(base kbEnt.Pattern) error {
	ex := ExtraDataPatternForRow{}
	ex.Fill(base.ExtraData)

	b, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	kr.ExtraData = string(b)
	return nil
}

func (pp ParameterPatternForRow) Extract() kbEnt.ParameterPattern {
	p := kbEnt.ParameterPattern{}
	p.ShortName = pp.ShortName
	p.Type = kbEnt.TypeParameter(pp.Type)
	return p
}

func (ex ExtraDataPatternForRow) Extract() kbEnt.ExtraDataPattern {
	e := kbEnt.ExtraDataPattern{}
	e.Description = ex.Description
	e.Language = kbEnt.Language(ex.Language)
	e.Script = ex.Script
	e.InputParameters = []kbEnt.ParameterPattern{}
	for _, v := range ex.InputParameters {
		e.InputParameters = append(e.InputParameters, v.Extract())
	}
	e.OutputParameters = []kbEnt.ParameterPattern{}
	for _, v := range ex.OutputParameters {
		e.OutputParameters = append(e.OutputParameters, v.Extract())
	}
	return e
}

func (pr PatternRow) Extract() (kbEnt.Pattern, error) {
	p := kbEnt.Pattern{}
	p.BaseInfo = pr.BaseInfoForRow.Extract()
	p.Type = kbEnt.TypePattern(pr.Type)

	ex, err := pr.extractExtraData()
	if err != nil {
		return kbEnt.Pattern{}, err
	}
	p.ExtraData = ex
	return p, nil
}

func (pr PatternRow) extractExtraData() (kbEnt.ExtraDataPattern, error) {
	ex := &ExtraDataPatternForRow{}
	err := json.Unmarshal([]byte(pr.ExtraData), &ex)
	if err != nil {
		return kbEnt.ExtraDataPattern{}, nil
	}
	return ex.Extract(), nil
}
