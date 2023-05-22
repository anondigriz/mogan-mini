package knowledgebases

import (
	"context"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

type Storage struct {
	lg  *zap.Logger
	fb  *filesbroker.FilesBroker
	c   *collector.Collector
	cfg config.Config
}

func New(lg *zap.Logger, cfg config.Config) *Storage {
	fb := filesbroker.New(lg, cfg.WorkspaceDir)
	c := collector.New(lg)
	st := &Storage{
		lg:  lg,
		fb:  fb,
		c:   c,
		cfg: cfg,
	}

	return st
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
