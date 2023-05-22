package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (st Storage) GetContainerByUUID(uuid string) (*kbEnt.Container, error) {
	filePath := st.fb.GetFilePath(uuid)
	return st.GetContainerByPath(filePath)
}

func (st Storage) GetContainerByPath(filePath string) (*kbEnt.Container, error) {
	uuid := st.fb.GetFileUUID(filePath)
	from, err := st.fb.OpenFileByPath(filePath)
	if err != nil {
		st.lg.Error(errMsgs.OpenKnowledgeBaseFileFail, zap.Error(err))
		return nil, err
	}
	defer from.Close()

	p := parser.New(st.lg)

	cont, err := p.Parse(parser.ParseXMLArgs{
		KnowledgeBaseUUID: uuid,
		XMLFile:           from,
		FileName:          from.Name(),
	})
	if err != nil {
		st.lg.Error(errMsgs.ParseKnowledgeBaseFail, zap.Error(err))
		return nil, errors.WrapExchangeErr(err)
	}
	return &cont, nil
}

func (st Storage) SaveContainer(cont *kbEnt.Container) error {
	filePath := st.fb.GetFilePath(cont.KnowledgeBase.UUID)
	to, err := st.fb.CreateFileByPath(filePath)
	if err != nil {
		st.lg.Error(errMsgs.CreateKnowledgeBaseFileFail, zap.Error(err))
		return err
	}
	defer to.Close()

	err = st.c.Collect(collector.ParseXMLArgs{
		Version: formats.VersionV3M0,
		Cont:    cont,
		XMLFile: to,
		Prefix:  st.cfg.XMLPrefix,
		Indent:  st.cfg.XMLIndent,
	})

	if err != nil {
		return err
	}
	return nil
}

func (st Storage) RemoveContainerByUUID(uuid string) error {
	return st.fb.RemoveFileByUUID(uuid)
}
