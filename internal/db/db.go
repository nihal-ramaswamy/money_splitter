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

func GetAllEmails(db *sql.DB) []string {
	if db == nil {
		panic("db cannot be nil")
	}

	query := `SELECT ID, EMAIL FROM "USER"`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var id string
		var email string

		err := rows.Scan(&id, &email)
		if err != nil {
			panic(err)
		}

		emails = append(emails, email)
	}

	return emails
}
