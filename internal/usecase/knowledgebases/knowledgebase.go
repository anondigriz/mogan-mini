package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
)

func (kb KnowledgeBases) RemoveKnowledgeBase(uuid string) error {
	return kb.kb.Remove(uuid)
}

func (kb KnowledgeBases) CreateKnowledgeBase(name string) (string, error) {
	return kb.kb.Create(name)
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
