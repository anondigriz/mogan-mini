package parser

import (
	"bufio"
	"context"
	"io"

	"github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/detector"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/exchange/v2m0"
)

const (
	versionV2M0 = "2.0"
)

type ParseXMLArgs struct {
	KnowledgeBaseUUID string
	XMLFile           argsCore.UploadedFile
	FileName          string
}

type Parser struct {
	lg       *zap.Logger
	cfg      config.Config
	detector *detector.Detector
	v2m0     *v2m0.V2M0
}

func New(lg *zap.Logger, cfg config.Config) *Parser {
	d := detector.New(lg)

	pv2m0 := v2m0.New(lg)
	i := &Parser{
		lg:       lg,
		cfg:      cfg,
		detector: d,
		v2m0:     pv2m0,
	}
	return i
}

func (p Parser) Parse(ctx context.Context, args ParseXMLArgs) (kbEnt.Container, error) {
	scanner := bufio.NewScanner(args.XMLFile)
	ver, err := p.detector.DetectVersion(scanner)
	if err != nil {
		p.lg.Error("xml exchange document file version could not be detected", zap.Error(err))
		return kbEnt.Container{}, errors.WrapXMLValidationErr(err)
	}

	if ver != versionV2M0 {
		p.lg.Error("unsupported format XML version", zap.Error(err))
		return kbEnt.Container{}, errors.NewUnsupportedFormatXMLVersionErr(ver)
	}

	err = p.seekFileToBegin(args)
	if err != nil {
		p.lg.Error("fail to seek file to the begin", zap.Error(err))
		return kbEnt.Container{}, err
	}

	return p.parseV2M0(args)
}

func (p Parser) seekFileToBegin(args ParseXMLArgs) error {
	_, err := args.XMLFile.Seek(0, 0)
	if err != nil {
		if err != nil {
			p.lg.Error("fail to reset the XML file reading stream to the beginning", zap.Error(err))
			return errors.NewReadingXMLFailErr(err)
		}
	}
	return nil
}

func (p Parser) parseV2M0(args ParseXMLArgs) (kbEnt.Container, error) {
	content, err := io.ReadAll(args.XMLFile)
	if err != nil {
		if err != nil {
			p.lg.Error("fail to read the XML file from stream", zap.Error(err))
			return kbEnt.Container{}, errors.NewReadingXMLFailErr(err)
		}
	}
	cont, err := p.v2m0.ParseXML(args.KnowledgeBaseUUID, content)
	if err != nil {
		if err != nil {
			p.lg.Error("fail to parse the XML file", zap.Error(err))
			return kbEnt.Container{}, errors.NewParsingXMLFailErr("fail to parse the XML file", err)
		}
	}
	return cont, nil
}
