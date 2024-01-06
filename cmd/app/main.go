package main

import (
	healthcheck_api "money_splitter/internal/api/health_check"
	db_config "money_splitter/internal/config/db"
	server_config "money_splitter/internal/config/server"

	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
)

func main() {
	db_config_default := db_config.Default()
	psqlIfo := db_config.GetPsqlInfo(db_config_default)

	server_config := server_config.Default()
	server := gin.Default()

	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", psqlIfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Pinging database...")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to server...")
	server.GET("/health_check", healthcheck_api.HealthCheckHandler())
	fx.New(
		fx.Invoke(
			server.Run(server_config.Port)),
	).Run()
}
