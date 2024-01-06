package serverconfig

type Config struct {
	Port string
}

func New(options ...func(*Config)) *Config {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return config
}

func WithPort(port string) func(*Config) {
	return func(c *Config) {
		c.Port = port
	}
}
