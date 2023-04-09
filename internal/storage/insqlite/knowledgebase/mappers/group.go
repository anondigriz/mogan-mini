package mappers

import (
	"github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type GroupRow struct {
	BaseInfoForRow
	KnowledgeBaseUUID string
	ExtraData         ExtraDataGroupForRow
}

type ExtraDataGroupForRow struct {
	Description string `json:"description"`
}

func (ex *ExtraDataGroupForRow) Fill(base knowledgebase.ExtraDataGroup) {
	ex.Description = base.Description
}

func (gr *GroupRow) Fill(base knowledgebase.Group) {
	gr.BaseInfoForRow.Fill(base.BaseInfo)
	gr.ExtraData.Fill(base.ExtraData)
}

func (ex *ExtraDataGroupForRow) Extract() knowledgebase.ExtraDataGroup {
	e := knowledgebase.ExtraDataGroup{}
	e.Description = ex.Description
	return e
}

func (gr *GroupRow) Extract() knowledgebase.Group {
	k := knowledgebase.Group{}
	k.BaseInfo = gr.BaseInfoForRow.Extract()
	k.ExtraData = gr.ExtraData.Extract()
	return k
}
