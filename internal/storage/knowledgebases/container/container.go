package container

import (
	"path"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
)

const (
	fileExtension = ".toml"
)

type Container struct {
	lg                *zap.Logger
	knowledgeBaseUUID string
	knowledgeBaseDir  string
}

func New(lg *zap.Logger, knowledgeBasesDir, knowledgeBaseUUID string) *Container {
	c := &Container{
		lg:               lg,
		knowledgeBaseDir: path.Join(knowledgeBasesDir, knowledgeBaseUUID),
	}
	return c
}

func (c Container) WriteContainer(cont *kbEnt.Container) error {
	if err := c.WriteKnowledgeBase(cont.KnowledgeBase); err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	if err := c.WriteGroups(cont.Groups); err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	if err := c.WriteParameters(cont.Parameters); err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	if err := c.WritePatterns(cont.Patterns); err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	if err := c.WriteRules(cont.Rules); err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	return nil
}

func (c Container) ReadContainer() (*kbEnt.Container, error) {
	cont := &kbEnt.Container{}
	kb, err := c.ReadKnowledgeBase(c.knowledgeBaseUUID)
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return nil, err
	}
	cont.KnowledgeBase = kb

	groups, err := c.ReadGroups()
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return nil, err
	}
	cont.Groups = groups

	parameters, err := c.ReadParameters()
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return nil, err
	}
	cont.Parameters = parameters

	patterns, err := c.ReadPatterns()
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return nil, err
	}
	cont.Patterns = patterns

	rules, err := c.ReadRules()
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return nil, err
	}
	cont.Rules = rules

	return cont, nil
}

func (c Container) RemoveContainer() error {
	fb := filesbroker.New(c.lg, c.knowledgeBaseDir, fileExtension)

	err := fb.RemoveDirByPath(c.knowledgeBaseDir)
	if err != nil {
		c.lg.Error(errMsgs.DeleteDirFail, zap.Error(err))
		return errors.NewDeleteDirFailErr(err, c.knowledgeBaseDir)
	}
	return nil
}
