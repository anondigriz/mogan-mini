package knowledgebase

import (
	"context"
	"os"

	"github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
)

type Storage struct {
	lg       *zap.Logger
	filePath string
	cont     *kbEnt.Container
}

func New(ctx context.Context, lg *zap.Logger, filePath string, uuid string) (*Storage, error) {
	from, err := os.Open(filePath)
	if err != nil {
		return nil, errors.NewUnexpectedStorageFailErr(err)
	}
	defer from.Close()

	p := parser.New(lg)

	cont, err := p.Parse(parser.ParseXMLArgs{
		KnowledgeBaseUUID: uuid,
		XMLFile:           from,
		FileName:          from.Name(),
	})
	if err != nil {
		// TODO log
		return nil, errors.NewUnexpectedStorageFailErr(err)
	}
	st := &Storage{
		lg:       lg,
		cont:     &cont,
		filePath: filePath,
	}

	return st, nil
}

func (st *Storage) GetPath() string {
	return st.filePath
}

func (st *Storage) Commit() error {
	w := newWriter(st.lg)
	return w.write(st.cont, st.filePath)
}

func (st *Storage) Shutdown() {
	st.cont = nil
}

func (st *Storage) Ping(ctx context.Context) error {
	return nil
}

func (st *Storage) GetKnowledgeBase(ctx context.Context) knowledgebase.KnowledgeBase {
	return st.cont.KnowledgeBase
}

func (st *Storage) UpdateKnowledgeBase(ctx context.Context, kb knowledgebase.KnowledgeBase) {
	st.cont.KnowledgeBase = kb
}

func (st *Storage) FillFromContainer(ctx context.Context, c kbEnt.Container) {
	st.cont = &c
}
