package container

import (
	"go.uber.org/zap"
)

const (
	fileExtension = ".toml"
)

type Container struct {
	lg               *zap.Logger
	knowledgeBaseDir string
}

func New(lg *zap.Logger, knowledgeBaseDir string) *Container {
	c := &Container{
		lg:               lg,
		knowledgeBaseDir: knowledgeBaseDir,
	}
	return c
}
