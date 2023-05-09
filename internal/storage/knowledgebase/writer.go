package knowledgebase

import (
	"os"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	"go.uber.org/zap"
)

type writer struct {
	lg *zap.Logger
}

func newWriter(lg *zap.Logger) *writer {
	w := &writer{
		lg: lg,
	}
	return w
}

func (w writer) write(cont *kbEnt.Container, filePath string) error {
	to, err := os.Create(filePath)
	if err != nil {
		// TODO log
		return errors.NewUnexpectedStorageFailErr(err)
	}
	defer to.Close()

	c := collector.New(w.lg)
	err = c.Collect(collector.ParseXMLArgs{
		Version: formats.VersionV3M0,
		Cont:    cont,
		XMLFile: to,
		Prefix:  "",
		Indent:  "  ",
	})

	if err != nil {
		return err
	}
	return nil
}
