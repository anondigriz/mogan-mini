package config

type CLICfg struct {
	Debug bool
}

func (c *CLICfg) setToCfg(cfg Config) Config {
	if c.Debug {
		cfg.Debug = c.Debug
	}
	return cfg
}
