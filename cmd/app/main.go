package main

import (
	healthcheck_api "money_splitter/internal/api/health_check"
	db_config "money_splitter/internal/config/db"
	server_config "money_splitter/internal/config/server"

	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
)

func main() {
	// Connection to database
	db_config := db_config.New(
		db_config.WithHost("host.docker.internal"),
		db_config.WithPort("5432"),
		db_config.WithUser("postgres"),
		db_config.WithPassword("postgres"),
		db_config.WithDbname("money_splitter"),
	)

	psqlIfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_config.Host, db_config.Port, db_config.User, db_config.Password, db_config.Dbname)

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

	// Connecting to server
	log.Println("Connecting to server...")
	server_config := server_config.New(
		server_config.WithPort(":8080"),
	)

	server := gin.Default()

	server.GET("/health_check", healthcheck_api.HealthCheckHandler())

	fx.New(
		fx.Invoke(
			server.Run(server_config.Port)),
	).Run()
}
