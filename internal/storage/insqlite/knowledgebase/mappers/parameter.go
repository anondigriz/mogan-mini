package mappers

import (
	"github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type ParameterRow struct {
	BaseInfoForRow
	GroupUUID   string
	Type      int
	ExtraData ExtraDataParameterForRow
}

type ExtraDataParameterForRow struct {
	Description  string `json:"description"`
	DefaultValue string `json:"defaultValue"`
}

func (ex *ExtraDataParameterForRow) Fill(base knowledgebase.ExtraDataParameter) {
	ex.Description = base.Description
	ex.DefaultValue = base.DefaultValue
}

func (pr *ParameterRow) Fill(base knowledgebase.Parameter) {
	pr.BaseInfoForRow.Fill(base.BaseInfo)
	pr.GroupUUID = base.GroupUUID
	pr.Type = int(base.Type)
	pr.ExtraData.Fill(base.ExtraData)
}

func (ex *ExtraDataParameterForRow) Extract() knowledgebase.ExtraDataParameter {
	e := knowledgebase.ExtraDataParameter{}
	e.Description = ex.Description
	e.DefaultValue = ex.DefaultValue
	return e
}

func (pr *ParameterRow) Extract() knowledgebase.Parameter {
	k := knowledgebase.Parameter{}
	k.BaseInfo = pr.BaseInfoForRow.Extract()
	k.GroupUUID = pr.GroupUUID
	k.Type = knowledgebase.TypeParameter(pr.Type)
	k.ExtraData = pr.ExtraData.Extract()
	return k
}
