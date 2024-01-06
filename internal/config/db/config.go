package db_config

import (
	"fmt"
	utils "money_splitter/internal/utils"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func New(options ...func(*Config)) *Config {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return config
}

func WithHost(host string) func(*Config) {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port string) func(*Config) {
	return func(c *Config) {
		c.Port = port
	}
}

func WithUser(user string) func(*Config) {
	return func(c *Config) {
		c.User = user
	}
}

func WithPassword(password string) func(*Config) {
	return func(c *Config) {
		c.Password = password
	}
}

func WithDbname(dbname string) func(*Config) {
	return func(c *Config) {
		c.Dbname = dbname
	}
}

func Default() *Config {
	return New(
		WithHost(utils.GetDotEnvVariable("DB_HOST")),
		WithPort(utils.GetDotEnvVariable("DB_PORT")),
		WithUser(utils.GetDotEnvVariable("DB_USER")),
		WithPassword(utils.GetDotEnvVariable("DB_PASSWORD")),
		WithDbname(utils.GetDotEnvVariable("DB_NAME")),
	)
}

func GetPsqlInfo(config *Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

}
