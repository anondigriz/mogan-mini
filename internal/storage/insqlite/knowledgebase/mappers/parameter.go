package mappers

import (
	"encoding/json"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type ParameterRow struct {
	BaseInfoForRow
	GroupUUID string
	Type      int
	ExtraData string
}

type ExtraDataParameterForRow struct {
	Description  string `json:"description"`
	DefaultValue string `json:"defaultValue"`
}

func (ex *ExtraDataParameterForRow) Fill(base kbEnt.ExtraDataParameter) {
	ex.Description = base.Description
	ex.DefaultValue = base.DefaultValue
}

func (pr *ParameterRow) Fill(base kbEnt.Parameter) error {
	pr.BaseInfoForRow.Fill(base.BaseInfo)
	pr.GroupUUID = base.GroupUUID
	pr.Type = int(base.Type)
	err := pr.fillExtraData(base)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ParameterRow) fillExtraData(base kbEnt.Parameter) error {
	ex := ExtraDataParameterForRow{}
	ex.Fill(base.ExtraData)
	b, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	pr.ExtraData = string(b)
	return nil
}

func (ex ExtraDataParameterForRow) Extract() kbEnt.ExtraDataParameter {
	e := kbEnt.ExtraDataParameter{}
	e.Description = ex.Description
	e.DefaultValue = ex.DefaultValue
	return e
}

func (pr ParameterRow) Extract() (kbEnt.Parameter, error) {
	p := kbEnt.Parameter{}
	p.BaseInfo = pr.BaseInfoForRow.Extract()
	p.GroupUUID = pr.GroupUUID
	p.Type = kbEnt.TypeParameter(pr.Type)

	ex, err := pr.extractExtraData()
	if err != nil {
		return kbEnt.Parameter{}, err
	}
	p.ExtraData = ex
	return p, nil
}

func (pr ParameterRow) extractExtraData() (kbEnt.ExtraDataParameter, error) {
	ex := &ExtraDataParameterForRow{}
	err := json.Unmarshal([]byte(pr.ExtraData), &ex)
	if err != nil {
		return kbEnt.ExtraDataParameter{}, nil
	}
	return ex.Extract(), nil
}
