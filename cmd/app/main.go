package main

import (
	healthcheck_api "money_splitter/internal/api/health_check"
	server_config "money_splitter/internal/config/server"
	db_instance "money_splitter/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	server_config := server_config.Default()
	server := gin.Default()

	db := db_instance.GetDbInstanceWithDefaultConfig()
	defer db.Close()

	server.GET("/health_check", healthcheck_api.HealthCheckHandler())
	server.Run(server_config.Port)
}
