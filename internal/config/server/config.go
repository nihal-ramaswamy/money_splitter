package serverconfig

import (
	utils "money_splitter/internal/utils"
)

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

func Default() *Config {
	return New(
		WithPort(":" + utils.GetDotEnvVariable("SERVER_PORT")),
	)
}
