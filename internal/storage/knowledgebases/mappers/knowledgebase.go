package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

type KnowledgeBase struct {
	BaseInfo
}

func (k *KnowledgeBase) Fill(base kbEnt.KnowledgeBase) {
	k.BaseInfo.Fill(base.BaseInfo)
}

func (k KnowledgeBase) Extract() kbEnt.KnowledgeBase {
	result := kbEnt.KnowledgeBase{
		BaseInfo: k.BaseInfo.Extract(),
	}
	return result
}
