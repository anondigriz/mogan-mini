package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

func (kb KnowledgeBases) RemoveGroup(knowledgeBaseUUID, uuid string) error {
	return kb.gr.Remove(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) CreateGroup(knowledgeBaseUUID string, ent kbEnt.Group) (string, error) {
	return kb.gr.Create(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetGroup(knowledgeBaseUUID, uuid string) (kbEnt.Group, error) {
	return kb.gr.Get(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) UpdateGroup(knowledgeBaseUUID string, ent kbEnt.Group) error {
	return kb.gr.Update(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetAllGroups(knowledgeBasesUUID string) (map[string]kbEnt.Group, error) {
	return kb.gr.GetAll(knowledgeBasesUUID)
}
