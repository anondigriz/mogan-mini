package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

func (kb KnowledgeBases) RemoveParameter(knowledgeBaseUUID, uuid string) error {
	return kb.par.Remove(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) CreateParameter(knowledgeBaseUUID string, ent kbEnt.Parameter) (string, error) {
	return kb.par.Create(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetParameter(knowledgeBaseUUID, uuid string) (kbEnt.Parameter, error) {
	return kb.par.Get(knowledgeBaseUUID, uuid)
}

func (kb KnowledgeBases) UpdateParameter(knowledgeBaseUUID string, ent kbEnt.Parameter) error {
	return kb.par.Update(knowledgeBaseUUID, ent)
}

func (kb KnowledgeBases) GetAllParameters(knowledgeBaseUUID string) (map[string]kbEnt.Parameter, error) {
	return kb.par.GetAll(knowledgeBaseUUID)
}
