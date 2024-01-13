package main

import (
	auth_api "money_splitter/internal/api/auth"
	group_api "money_splitter/internal/api/group"
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

	// Postgres DB
	pdb := db_instance.GetPostgresDbInstanceWithDefaultConfig()
	defer pdb.Close()

	// Redis DB
	rdb := db_instance.GetRedisDbInstanceWithDefaultConfig()

	// Routes
	server.GET("/health_check", healthcheck_api.HealthCheckHandler())
	server.GET("/health_check_auth", middleware.AuthMiddleware(pdb, rdb), healthcheck_api.HealthCheckHandlerAuth())

	auth := server.Group("/auth")
	{
		auth.POST("/register", auth_api.NewUserHandler(pdb))
		auth.POST("/login", auth_api.LoginUserHandler(pdb, rdb))
	}

	auth2 := server.Group("/auth", middleware.AuthMiddleware(pdb, rdb))
	{
		auth2.POST("/logout", auth_api.LogoutUserHandler(rdb))
	}

	group := server.Group("/group", middleware.AuthMiddleware(pdb, rdb))
	{
		group.POST("/newGroup", group_api.NewGroupHandler(pdb, rdb))
	}

	server.Run(server_config.Port)
}
