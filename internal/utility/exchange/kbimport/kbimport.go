package kbimport

import (
	"bufio"
	"context"
	"io"

	"github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/validator"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/core/args"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/utility/errors"
	"github.com/anondigriz/mogan-mini/internal/utility/exchange/kbimport/v2m0"
	"go.uber.org/zap"
)

const (
	versionV2M0 = "2.0"
)

type KBImport struct {
	lg        *zap.Logger
	cfg       config.Config
	validator *validator.Validator
	v2m0      *v2m0.V2M0
}

func New(lg *zap.Logger, cfg config.Config) (*KBImport, error) {
	v, err := validator.New(lg)
	if err != nil {
		lg.Error("xml validator initialization error", zap.Error(err))
		return nil, err
	}
	parser := v2m0.New(lg)
	i := &KBImport{
		lg:        lg,
		cfg:       cfg,
		validator: v,
		v2m0:      parser,
	}
	return i, nil
}

func (kb *KBImport) Parse(ctx context.Context, arg args.ImportKnowledgeBase) (kbEnt.Container, error) {
	scanner := bufio.NewScanner(arg.XMLFile)
	ver, err := kb.validator.DetectVersion(scanner)
	if err != nil {
		kb.lg.Error("xml exchange document file version could not be detected", zap.Error(err))
		return kbEnt.Container{}, errors.WrapXMLValidationErr(err)
	}

	if ver != versionV2M0 {
		kb.lg.Error("unsupported format XML version", zap.Error(err))
		return kbEnt.Container{}, errors.NewUnsupportedFormatXMLVersionErr(ver)
	}

	err = kb.seekFileToBegin(arg)
	if err != nil {
		kb.lg.Error("fail to seek file to the begin", zap.Error(err))
		return kbEnt.Container{}, err
	}
	return kb.parseV2M0(arg)
}

func (kb *KBImport) seekFileToBegin(arg args.ImportKnowledgeBase) error {
	_, err := arg.XMLFile.Seek(0, 0)
	if err != nil {
		if err != nil {
			kb.lg.Error("fail to reset the XML file reading stream to the beginning", zap.Error(err))
			return errors.NewReadingXMLFailErr(err)
		}
	}
	return nil
}

func (kb *KBImport) parseV2M0(arg args.ImportKnowledgeBase) (kbEnt.Container, error) {
	content, err := io.ReadAll(arg.XMLFile)
	if err != nil {
		if err != nil {
			kb.lg.Error("fail to read the XML file from stream", zap.Error(err))
			return kbEnt.Container{}, errors.NewReadingXMLFailErr(err)
		}
	}
	cont, err := kb.v2m0.ParseXML(arg.KnowledgeBaseUUID, content)
	if err != nil {
		if err != nil {
			kb.lg.Error("fail to parse the XML file", zap.Error(err))
			return kbEnt.Container{}, errors.NewParsingXMLFailErr("fail to parse the XML file", err)
		}
	}
	return cont, nil
}
