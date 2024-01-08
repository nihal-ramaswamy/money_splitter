package auth_api

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
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

		if err := c.ShouldBindJSON(&user); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
		}

		if db_utils.DoesEmailExist(db, user.Email) {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("User with email %s already exists", user.Email)})

		}

		id := db_utils.RegisterNewUser(db, user)

		c.JSON(http.StatusAccepted, gin.H{"id": id})

	}
}

func LoginUserHandler(db *sql.DB) func(c *gin.Context) {
	if db == nil {
		log.Panic("db cannot be nil")
	}
	return func(c *gin.Context) {
		var user dto.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if !db_utils.DoesEmailExist(db, user.Email) {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": fmt.Sprintf("User with email %s does not exist", user.Email)})
			return
		}

		if !db_utils.DoesPasswordMatch(db, user) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token, err := db_utils.GenerateToken(user)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}
