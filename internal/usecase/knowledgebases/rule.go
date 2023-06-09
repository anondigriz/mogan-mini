package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

func (kb KnowledgeBases) RemoveRule(knowledgeBaseUUID, uuid string) error {
	return kb.rul.Remove(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) CreateRule(knowledgeBaseUUID string, ent kbEnt.Rule) (string, error) {
	return kb.rul.Create(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetRule(knowledgeBaseUUID, uuid string) (kbEnt.Rule, error) {
	return kb.rul.Get(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) UpdateRule(knowledgeBaseUUID string, ent kbEnt.Rule) error {
	return kb.rul.Update(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetAllRules(knowledgeBaseUUID string) (map[string]kbEnt.Rule, error) {
	return kb.rul.GetAll(knowledgeBaseUUID)
}
