package main

import (
	auth_api "money_splitter/internal/api/auth"
	healthcheck_api "money_splitter/internal/api/health_check"
	server_config "money_splitter/internal/config/server"
	db_instance "money_splitter/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	server_config := server_config.Default()
	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	db := db_instance.GetDbInstanceWithDefaultConfig()
	defer db.Close()

	server.GET("/health_check", healthcheck_api.HealthCheckHandler())

	auth := server.Group("/auth")
	{
		auth.POST("/register", auth_api.NewUserHandler(db))
	}

	server.Run(server_config.Port)
}
