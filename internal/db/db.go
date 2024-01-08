package db

import (
	"database/sql"
	"log"
	db_config "money_splitter/internal/config/db"
	dto "money_splitter/internal/dto"
	"money_splitter/internal/utils"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

func selectAllWhereEmailIs(db *sql.DB, email string) (dto.User, error) {
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

func DoesEmailExist(db *sql.DB, email string) bool {
	_, err := selectAllWhereEmailIs(db, email)
	return err != sql.ErrNoRows
}

func RegisterNewUser(db *sql.DB, user dto.User) string {

	var id string
	user = user.HashAndSalt()

	err := db.QueryRow("INSERT INTO \"USER\" (NAME, EMAIL, PASSWORD, IS_VERIFIED) VALUES ($1, $2, $3, $4) RETURNING ID",
		user.Name, user.Email, user.Password, false).Scan(&id)

	if err != nil {
		log.Panic(err)
	}

	return id
}

func DoesPasswordMatch(db *sql.DB, user dto.User) bool {
	if db == nil {
		log.Panic("db cannot be nil")
	}

	var password string
	query := `SELECT PASSWORD FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, user.Email).Scan(&password)

	if err != nil {
		log.Panic(err)
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)) == nil
}

func GenerateToken(user dto.User) (string, error) {
	secret := utils.GetDotEnvVariable("SECRET_KEY")

	var signingKey = []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 60 * 60 * 24 * 365).Unix()

	return token.SignedString(signingKey)
}
