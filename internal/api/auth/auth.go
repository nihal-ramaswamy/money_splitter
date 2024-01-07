package auth_api

import (
	"database/sql"
	"fmt"
	"log"
	db_utils "money_splitter/internal/db"
	dto "money_splitter/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(db *sql.DB) func(c *gin.Context) {
	if db == nil {
		log.Panic("db cannot be nil")
	}

	return func(c *gin.Context) {
		var user dto.User
		var id string

		if err := c.ShouldBindJSON(&user); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
		}

		emails := db_utils.GetAllEmails(db)

		for _, email := range emails {
			if email == user.Email {
				c.JSON(http.StatusBadRequest,
					gin.H{"error": fmt.Sprintf("User with email %s already exists", user.Email)})
			}
		}

		user = user.HashAndSalt()

		err := db.QueryRow("INSERT INTO \"USER\" (NAME, EMAIL, PASSWORD, IS_VERIFIED) VALUES ($1, $2, $3, $4) RETURNING ID",
			user.Name, user.Email, user.Password, false).Scan(&id)

		if err != nil {
			log.Panic(err)
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Inserted user with id " + id + " successfully"})

	}
}
