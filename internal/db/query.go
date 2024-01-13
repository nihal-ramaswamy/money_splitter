package db

import (
	"database/sql"
	"money_splitter/internal/dto"
)

func selectAllFromUserWhereEmailIs(db *sql.DB, email string) (dto.User, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var user dto.User
	query := `SELECT * FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&user)

	if err != nil {
		return user, err
	}

	return user, err
}

func selectPasswordFromUserWhereEmailIDs(db *sql.DB, email string) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}
	var password string
	query := `SELECT PASSWORD FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&password)

	return password, err
}

// Returns ID of newly created user
func insertIntoUser(db *sql.DB, user dto.User) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var id string
	query := `INSERT INTO "USER" (NAME, EMAIL, PASSWORD, IS_VERIFIED) VALUES ($1, $2, $3, $4) RETURNING ID`
	err := db.QueryRow(query, user.Name, user.Email, user.Password, false).Scan(&id)

	return id, err
}

// Returns ID of newly created group
func insertIntoGroup(db *sql.DB, group dto.Group) (string, error) {
	if db == nil {
		panic("db cannot be null")
	}

	var id string
	query := `INSERT INTO "GROUP" (GROUP_NAME, GROUP_ADMIN, SIMPLIFY_TXN) VALUES ($1, $2, $3) RETURNING ID`
	err := db.QueryRow(query, group.GroupName, group.GroupAdmin, group.SimplifyTxn).Scan(&id)

	return id, err
}
