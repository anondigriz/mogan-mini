package mappers

import (
	"encoding/json"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type GroupRow struct {
	BaseInfoForRow
	ExtraData string
}

type ExtraDataGroupForRow struct {
	Description string `json:"description"`
}

func (ex *ExtraDataGroupForRow) Fill(base kbEnt.ExtraDataGroup) {
	ex.Description = base.Description
}

func (gr *GroupRow) Fill(base kbEnt.Group) error {
	gr.BaseInfoForRow.Fill(base.BaseInfo)
	err := gr.fillExtraData(base)
	if err != nil {
		return err
	}
	return nil
}

func (gr *GroupRow) fillExtraData(base kbEnt.Group) error {
	ex := ExtraDataGroupForRow{}
	ex.Fill(base.ExtraData)

	b, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	gr.ExtraData = string(b)
	return nil
}

func (ex ExtraDataGroupForRow) Extract() kbEnt.ExtraDataGroup {
	e := kbEnt.ExtraDataGroup{}
	e.Description = ex.Description
	return e
}

func (gr GroupRow) Extract() (kbEnt.Group, error) {
	g := kbEnt.Group{}
	g.BaseInfo = gr.BaseInfoForRow.Extract()

	ex, err := gr.extractExtraData()
	if err != nil {
		return kbEnt.Group{}, err
	}
	g.ExtraData = ex
	return g, nil
}

func (gr GroupRow) extractExtraData() (kbEnt.ExtraDataGroup, error) {
	ex := &ExtraDataGroupForRow{}
	err := json.Unmarshal([]byte(gr.ExtraData), &ex)
	if err != nil {
		return kbEnt.ExtraDataGroup{}, nil
	}
	return ex.Extract(), nil
}
