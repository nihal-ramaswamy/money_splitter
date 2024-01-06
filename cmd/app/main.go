package main

import (
	healthcheck_api "money_splitter/internal/api/health_check"
	db_config "money_splitter/internal/config/db"
	server_config "money_splitter/internal/config/server"

	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
	"os"
)

func getDotEnvVariable(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("key: ", key, "value: ", os.Getenv(key))

	return os.Getenv(key)
}

func main() {
	// Connection to database
	db_config := db_config.New(
		db_config.WithHost(getDotEnvVariable("DB_HOST")),
		db_config.WithPort(getDotEnvVariable("DB_PORT")),
		db_config.WithUser(getDotEnvVariable("DB_USER")),
		db_config.WithPassword(getDotEnvVariable("DB_PASSWORD")),
		db_config.WithDbname(getDotEnvVariable("DB_NAME")),
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
		server_config.WithPort(":" + getDotEnvVariable("SERVER_PORT")),
	)

	server := gin.Default()

	server.GET("/health_check", healthcheck_api.HealthCheckHandler())

	fx.New(
		fx.Invoke(
			server.Run(server_config.Port)),
	).Run()
}
