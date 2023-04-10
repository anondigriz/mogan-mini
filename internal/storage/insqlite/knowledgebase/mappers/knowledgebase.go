package mappers

import (
	"encoding/json"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
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

func (gr *GroupHierarchyForRow) Fill(base kbEnt.GroupHierarchy) {
	gr.GroupUUID = base.GroupUUID
	gr.Contains = []GroupHierarchyForRow{}
	for _, v := range base.Contains {
		e := GroupHierarchyForRow{}
		e.Fill(v)
		gr.Contains = append(gr.Contains, e)
	}
}

func (ex *ExtraDataKnowledgeBaseForRow) Fill(base kbEnt.ExtraDataKnowledgeBase) {
	ex.Description = base.Description
	ex.Groups.Fill(base.Groups)
}

func (kr *KnowledgeBaseRow) Fill(base kbEnt.KnowledgeBase) error {
	kr.BaseInfoForRow.Fill(base.BaseInfo)
	err := kr.fillExtraData(base)
	if err != nil {
		return err
	}
	return nil
}

func (kr *KnowledgeBaseRow) fillExtraData(base kbEnt.KnowledgeBase) error {
	ex := ExtraDataKnowledgeBaseForRow{}
	ex.Fill(base.ExtraData)

	b, err := json.Marshal(ex)
	if err != nil {
		return err
	}

	kr.ExtraData = string(b)
	return nil
}

func (gr GroupHierarchyForRow) Extract() kbEnt.GroupHierarchy {
	e := kbEnt.GroupHierarchy{GroupUUID: gr.GroupUUID}
	e.Contains = []kbEnt.GroupHierarchy{}
	for _, v := range gr.Contains {
		e.Contains = append(e.Contains, v.Extract())
	}

	return e
}

func (ex ExtraDataKnowledgeBaseForRow) Extract() kbEnt.ExtraDataKnowledgeBase {
	e := kbEnt.ExtraDataKnowledgeBase{}
	e.Description = ex.Description
	e.Groups = ex.Groups.Extract()
	return e
}

func (kr KnowledgeBaseRow) Extract() (kbEnt.KnowledgeBase, error) {
	k := kbEnt.KnowledgeBase{}
	k.BaseInfo = kr.BaseInfoForRow.Extract()

	ex, err := kr.extractExtraData()
	if err != nil {
		return kbEnt.KnowledgeBase{}, err
	}
	k.ExtraData = ex
	return k, nil
}

func (kr KnowledgeBaseRow) extractExtraData() (kbEnt.ExtraDataKnowledgeBase, error) {
	ex := &ExtraDataKnowledgeBaseForRow{}
	err := json.Unmarshal([]byte(kr.ExtraData), &ex)
	if err != nil {
		return kbEnt.ExtraDataKnowledgeBase{}, nil
	}
	return ex.Extract(), nil
}
