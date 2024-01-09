package db_config_redis

import (
	"money_splitter/internal/constants"
	"money_splitter/internal/utils"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func New(options ...func(*Config)) *Config {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return config
}

func WithAddr(addr string) func(*Config) {
	return func(c *Config) {
		c.Addr = addr
	}
}

func WithPassword(password string) func(*Config) {
	return func(c *Config) {
		c.Password = password
	}
}

func WithDB(db int) func(*Config) {
	return func(c *Config) {
		c.DB = db
	}
}

func Default() *Config {
	return New(
		WithAddr(utils.GetDotEnvVariable(constants.REDIS_HOST)+":"+utils.GetDotEnvVariable(constants.REDIS_PORT)),
		WithPassword(utils.GetDotEnvVariable(constants.REDIS_PASSWORD)),
		WithDB(0),
	)
}
