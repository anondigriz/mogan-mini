package mappers

import (
	"encoding/json"

	"github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
)

type KnowledgeBaseRow struct {
	BaseInfoForRow
	ExtraData string
}

type ExtraDataKnowledgeBaseForRow struct {
	Description string               `json:"description"`
	Groups      GroupHierarchyForRow `json:"groups"`
}

type GroupHierarchyForRow struct {
	GroupUUID string                 `json:"groupUUID"`
	Contains  []GroupHierarchyForRow `json:"contains"`
}

func (gr *GroupHierarchyForRow) Fill(base knowledgebase.GroupHierarchy) {
	gr.GroupUUID = base.GroupUUID
	gr.Contains = []GroupHierarchyForRow{}
	for _, v := range base.Contains {
		e := GroupHierarchyForRow{}
		e.Fill(v)
		gr.Contains = append(gr.Contains, e)
	}
}

func (ex *ExtraDataKnowledgeBaseForRow) Fill(base knowledgebase.ExtraDataKnowledgeBase) {
	ex.Description = base.Description
	ex.Groups.Fill(base.Groups)
}

func (kr *KnowledgeBaseRow) Fill(base knowledgebase.KnowledgeBase) error {
	kr.BaseInfoForRow.Fill(base.BaseInfo)

	ex := ExtraDataKnowledgeBaseForRow{}
	ex.Fill(base.ExtraData)

	b, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	kr.ExtraData = string(b)
	return err
}

func (gr *GroupHierarchyForRow) Extract() knowledgebase.GroupHierarchy {
	e := knowledgebase.GroupHierarchy{GroupUUID: gr.GroupUUID}
	e.Contains = []knowledgebase.GroupHierarchy{}
	for _, v := range gr.Contains {
		e.Contains = append(e.Contains, v.Extract())
	}

	return e
}

func (ex *ExtraDataKnowledgeBaseForRow) Extract() knowledgebase.ExtraDataKnowledgeBase {
	e := knowledgebase.ExtraDataKnowledgeBase{}
	e.Description = ex.Description
	e.Groups = ex.Groups.Extract()
	return e
}

func (kr *KnowledgeBaseRow) Extract() (knowledgebase.KnowledgeBase, error) {
	k := knowledgebase.KnowledgeBase{}
	k.BaseInfo = kr.BaseInfoForRow.Extract()

	ex := &ExtraDataKnowledgeBaseForRow{}
	err := json.Unmarshal([]byte(kr.ExtraData), &ex)
	if err != nil {
		return knowledgebase.KnowledgeBase{}, nil
	}

	k.ExtraData = ex.Extract()
	return k, nil
}
