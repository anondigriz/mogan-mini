package knowledgebase

import (
	"github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"
)

type StorageCreator struct {
	lg *zap.Logger
}

func NewStorageCreator(lg *zap.Logger) *StorageCreator {
	sc := &StorageCreator{
		lg: lg,
	}
	return sc
}

func (sc *StorageCreator) CreateStorage(kb knowledgebase.KnowledgeBase, filePath string) error {
	cont := &kbEnt.Container{
		KnowledgeBase: kb,
		Groups:        make(map[string]kbEnt.Group),
		Parameters:    make(map[string]kbEnt.Parameter),
		Patterns:      make(map[string]kbEnt.Pattern),
		Rules:         make(map[string]kbEnt.Rule),
	}

	w := newWriter(sc.lg)
	return w.write(cont, filePath)
}
