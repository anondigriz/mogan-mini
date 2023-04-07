package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Config struct {
		Test     string `mapstructure:"test"`
		Projects string
	}
)

func New(lg *zap.Logger, vp *viper.Viper, projects string) (Config, error) {
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

	err := vp.Unmarshal(&cfg)
	if err != nil {
		lg.Error("unable to decode into struct from file", zap.Error(err))
		return Config{}, err
	}
	lg.Debug("configuration has been set", zap.Reflect("config", cfg))

	return cfg, nil
}
