package knowledgebases

import (
	"context"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

type Settings struct {
	WorkspaceDir string
	XMLPrefix    string
	XMLIndent    string
}

type Storage struct {
	lg      *zap.Logger
	fb      *filesbroker.FilesBroker
	c       *collector.Collector
	setting Settings
}

func New(lg *zap.Logger, settings Settings) (*Storage, error) {
	fb := filesbroker.New(lg, settings.WorkspaceDir)
	c := collector.New(lg)
	st := &Storage{
		lg:      lg,
		fb:      fb,
		c:       c,
		setting: settings,
	}

	return st, nil
}

func (st Storage) Shutdown() {

}

func (st Storage) Ping(ctx context.Context) error {
	return nil
}

func (st Storage) GetAllKnowledgeBases() []kbEnt.KnowledgeBase {
	paths := st.fb.GetAllFilesPaths()

	kbs := []kbEnt.KnowledgeBase{}
	for _, filePath := range paths {
		kb, err := st.GetKnowledgeBase(st.fb.GetFileUUID(filePath))
		if err != nil {
			st.lg.Error(errMsgs.GetKnowledgeFail, zap.Error(err))
			continue
		}
		kbs = append(kbs, kb)
	}

	return kbs
}
