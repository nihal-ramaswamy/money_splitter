package db

import (
	"context"
	"database/sql"
	"log"
	db_config_postgres "money_splitter/internal/config/db/postgres"
	db_config_redis "money_splitter/internal/config/db/redis"
	dto "money_splitter/internal/dto"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var Ctx = context.Background()

func GetPostgresDbInstanceWithDefaultConfig() *sql.DB {
	psqlInfo := db_config_postgres.GetPsqlInfoDefault()

	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Pinging postgres...")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetRedisDbInstanceWithDefaultConfig() *redis.Client {
	redisInfo := db_config_redis.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisInfo.Addr,
		Password: redisInfo.Password,
		DB:       redisInfo.DB,
	})

	log.Println("Pinging redis...")
	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatal(err)
	}

	return rdb
}

func DoesEmailExist(db *sql.DB, email string) bool {
	_, err := selectAllFromUserWhereEmailIs(db, email)
	return err != sql.ErrNoRows
}

func GetUserFromEmail(db *sql.DB, email string) (dto.User, error) {
	return selectAllFromUserWhereEmailIs(db, email)
}

func RegisterNewUser(db *sql.DB, user dto.User) string {

	var id string
	user = user.HashAndSalt()

	id, err := insertIntoUser(db, user)

	if err != nil {
		log.Panic(err)
	}

	return id
}

func DoesPasswordMatch(db *sql.DB, user dto.User) bool {
	if db == nil {
		log.Panic("db cannot be nil")
	}

	password, err := selectPasswordFromUserWhereEmailIDs(db, user.Email)

	if err != nil {
		log.Panic(err)
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)) == nil
}

func RegisterNewGroup(db *sql.DB, group dto.Group) string {
	var id string
	id, err := insertIntoGroup(db, group)

	if err != nil {
		log.Panic(err)
	}

	return id
}
