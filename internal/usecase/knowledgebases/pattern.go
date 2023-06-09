package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

func (kb KnowledgeBases) RemovePattern(knowledgeBaseUUID, uuid string) error {
	return kb.pat.Remove(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) CreatePattern(knowledgeBaseUUID string, ent kbEnt.Pattern) (string, error) {
	return kb.pat.Create(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetPattern(knowledgeBaseUUID, uuid string) (kbEnt.Pattern, error) {
	return kb.pat.Get(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) UpdatePattern(knowledgeBaseUUID string, ent kbEnt.Pattern) error {
	return kb.pat.Update(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetAllPatterns(knowledgeBaseUUID string) (map[string]kbEnt.Pattern, error) {
	return kb.pat.GetAll(knowledgeBaseUUID)
}
