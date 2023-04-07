package config

import (
	"github.com/anondigriz/mogan-core/pkg/loglevel"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Databases struct {
		LogLevel          loglevel.LogLevel
		KnowledgeBasesDSN string `mapstructure:"KnowledgeBasesDSN"`
	}
	Config struct {
		// Test     string `mapstructure:"test"`
		Projects  string    `mapstructure:"-"`
		Databases Databases `mapstructure:"Databases"`
	}
)

func New(lg *zap.Logger, vp *viper.Viper, debug bool, projects string) (Config, error) {
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			lg.Error("config file was not found", zap.Error(err))
			return Config{}, err
		} else {
			lg.Error("fail to read config", zap.Error(err))
			return Config{}, err
		}
	}
	cfg := Config{
		Projects: projects,
	}
	cfg.setLogLevel(debug)

	err := vp.Unmarshal(&cfg)
	if err != nil {
		lg.Error("unable to decode into struct from file", zap.Error(err))
		return Config{}, err
	}
	lg.Debug("configuration has been set", zap.Reflect("config", cfg))

	return cfg, nil
}

func (c *Config) setLogLevel(debug bool) {
	if debug {
		c.Databases.LogLevel = loglevel.Debug
	} else {
		c.Databases.LogLevel = loglevel.Error
	}
}
