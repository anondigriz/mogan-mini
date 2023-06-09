package config

import (
	"github.com/anondigriz/mogan-core/pkg/loglevel"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Databases struct {
		LogLevel loglevel.LogLevel
	}
	KnowledgeBase struct {
		UUID string `mapstructure:"uuid"`
	}
	Config struct {
		// Test     string `mapstructure:"test"`
		IsDebug              bool
		WorkspaceDir         string        `mapstructure:"-"`
		Databases            Databases     `mapstructure:"databases"`
		CurrentKnowledgeBase KnowledgeBase `mapstructure:"currentknowledgebase"`
		XMLPrefix            string        `mapstructure:"xmlprefix"`
		XMLIndent            string        `mapstructure:"xmlindent"`
	}
)

func (c *Config) Fill(lg *zap.Logger, vp *viper.Viper, debug bool, workspaceDir string) error {
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			lg.Error("config file was not found", zap.Error(err))
			return err
		} else {
			lg.Error("fail to read config", zap.Error(err))
			return err
		}
	}
	c.WorkspaceDir = workspaceDir
	c.setLogLevel(debug)
	c.IsDebug = debug

	err := vp.Unmarshal(c)
	if err != nil {
		lg.Error("unable to decode into struct from file", zap.Error(err))
		return err
	}
	lg.Debug("configuration has been set", zap.Reflect("config", c))

	// TODO Make management of the following settings
	c.XMLPrefix = ""
	c.XMLIndent = "  "

	return nil
}

func (c *Config) setLogLevel(debug bool) {
	if debug {
		c.Databases.LogLevel = loglevel.Debug
	} else {
		c.Databases.LogLevel = loglevel.Error
	}
}
