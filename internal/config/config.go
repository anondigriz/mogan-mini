package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Config struct {
		Debug       bool   `mapstructure:"debug"`
		Projectbase string `mapstructure:"projectbase"`
		Test        string `mapstructure:"test"`
	}
)

func New(lg *zap.Logger, vp *viper.Viper) (Config, error) {
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			lg.Error("fail to read config", zap.Error(err))
			return Config{}, err
		}
	}
	cfg := Config{}

	err := vp.Unmarshal(&cfg)
	if err != nil {
		lg.Error("unable to decode into struct from file", zap.Error(err))
		return Config{}, err
	}
	lg.Info("configuration has been set", zap.Reflect("config", cfg))
	return cfg, nil
}
