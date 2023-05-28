package knowledgebases

import (
	"context"
	"path"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

const (
	KnowledgeBasesSubDir = "knowledgebases"
)

type Storage struct {
	lg                   *zap.Logger
	KnowledgeBasesSubDir string
}

func New(lg *zap.Logger, workspaceDir string) *Storage {
	st := &Storage{
		lg:                   lg,
		KnowledgeBasesSubDir: path.Join(workspaceDir, KnowledgeBasesSubDir),
	}

	return st
}

func (st Storage) Shutdown() {

}

func (st Storage) Ping(ctx context.Context) error {
	return nil
}

func (st Storage) GetAllKnowledgeBases() []kbEnt.KnowledgeBase {
	fb := filesbroker.New(st.lg, KnowledgeBasesSubDir, "")
	paths := fb.GetAllChildDirNames()
	kbs := []kbEnt.KnowledgeBase{}
	for _, name := range paths {
		kb, err := st.GetKnowledgeBase(name)
		if err != nil {
			st.lg.Error(errMsgs.GetKnowledgeFail, zap.Error(err))
			continue
		}
		kbs = append(kbs, kb)
	}

	return kbs
}
