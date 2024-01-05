package main

import (
	healthcheck_api "money_splitter/internal/api/health_check"
	server_config "money_splitter/internal/server_config"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	config := server_config.New(
		server_config.WithPort(":8080"),
	)

	server := gin.Default()

	server.GET("/health_check", healthcheck_api.HealthCheckHandler())

	fx.New(
		fx.Invoke(
			server.Run(config.Port)),
	).Run()
}
