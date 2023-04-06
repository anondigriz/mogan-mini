package config

type (
	Config struct {
		Debug bool
	}
)

// New returns app config.
func New(cli CLICfg) (Config, error) {
	cfg := Config{}
	cfg = cli.setToCfg(cfg)

	return cfg, nil
}
