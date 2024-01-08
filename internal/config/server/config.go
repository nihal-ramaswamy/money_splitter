package serverconfig

import (
	utils "money_splitter/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port    string
	GinMode string
	Cors    cors.Config
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

func WithGinMode(mode string) func(*Config) {
	var ginMode string
	switch mode {
	case "dev":
		ginMode = gin.DebugMode
	case "prod":
		ginMode = gin.ReleaseMode
	case "test":
		ginMode = gin.TestMode
	}

	return func(c *Config) {
		c.GinMode = ginMode
	}
}

func WithCors(config cors.Config) func(*Config) {
	return func(c *Config) {
		c.Cors = config
	}
}

func WithCorsHosts(hosts []string) func(*Config) {
	return func(c *Config) {
		c.Cors.AllowOrigins = hosts
	}
}

func Default() *Config {
	return New(
		WithPort(utils.GetDotEnvVariable("SERVER_PORT")),
		WithGinMode(utils.GetDotEnvVariable("ENV")),
		WithCors(cors.DefaultConfig()),
		WithCorsHosts([]string{utils.GetDotEnvVariable("SERVER_HOST") + utils.GetDotEnvVariable("SERVER_PORT")}),
	)
}
