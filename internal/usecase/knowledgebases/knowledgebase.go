package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"

	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
)

func (kb KnowledgeBases) RemoveKnowledgeBase(uuid string) error {
	return kb.kb.Remove(uuid)
}

func (kb KnowledgeBases) CreateKnowledgeBase(ent kbEnt.KnowledgeBase) (string, error) {
	return kb.kb.Create(ent)
}

func (kb KnowledgeBases) ImportKnowledgeBase(args argsCore.ImportKnowledgeBase) (string, error) {
	return kb.kb.Import(args)
}

func (kb KnowledgeBases) GetKnowledgeBase(uuid string) (kbEnt.KnowledgeBase, error) {
	return kb.kb.Get(uuid)
}

func (kb KnowledgeBases) UpdateKnowledgeBase(ent kbEnt.KnowledgeBase) error {
	return kb.kb.Update(ent)
}

func (kb KnowledgeBases) GetAllKnowledgeBases() map[string]kbEnt.KnowledgeBase {
	return kb.kb.GetAll()
}
