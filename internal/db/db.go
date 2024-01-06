package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	db_config "money_splitter/internal/config/db"
)

func GetDbInstanceWithDefaultConfig() *sql.DB {
	psqlIfo := db_config.GetPsqlInfoDefault()

	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", psqlIfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Pinging database...")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
