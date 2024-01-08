package main

import (
	auth_api "money_splitter/internal/api/auth"
	healthcheck_api "money_splitter/internal/api/health_check"
	"money_splitter/internal/api/middleware"
	server_config "money_splitter/internal/config/server"
	db_instance "money_splitter/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server_config := server_config.Default()

	gin.SetMode(server_config.GinMode)
	server := gin.Default()
	server.Use(cors.New(server_config.Cors))

	// DB
	db := db_instance.GetDbInstanceWithDefaultConfig()
	defer db.Close()

	// Routes
	server.GET("/health_check", healthcheck_api.HealthCheckHandler())
	server.GET("/health_check_auth", middleware.AuthMiddleware(), healthcheck_api.HealthCheckHandlerAuth())

	auth := server.Group("/auth")
	{
		auth.POST("/register", auth_api.NewUserHandler(db))
		auth.POST("/login", auth_api.LoginUserHandler(db))
	}

	server.Run(server_config.Port)
}
