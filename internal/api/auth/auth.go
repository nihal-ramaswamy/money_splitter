package auth_api

import (
	"database/sql"
	"fmt"
	"log"
	"money_splitter/internal/constants"
	db_utils "money_splitter/internal/db"
	dto "money_splitter/internal/dto"
	"money_splitter/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

func LoginUserHandler(pdb *sql.DB, rdb *redis.Client) func(c *gin.Context) {
	if pdb == nil || rdb == nil {
		log.Panic("db cannot be nil")
	}
	return func(c *gin.Context) {
		var user dto.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if !db_utils.DoesEmailExist(pdb, user.Email) {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": fmt.Sprintf("User with email %s does not exist", user.Email)})
			return
		}

		if !db_utils.DoesPasswordMatch(pdb, user) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(user)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.SetCookie(
			user.Email,
			token,
			int(constants.TOKEN_EXPIRY_TIME),
			"/",
			utils.GetDotEnvVariable("SERVER_HOST"),
			false,
			true,
		)

		rdb.Set(db_utils.Ctx, user.Email, token, constants.TOKEN_EXPIRY_TIME)

		c.JSON(http.StatusAccepted, gin.H{"message": "ok"})
	}
}

func LogoutUserHandler(rdb *redis.Client) func(c *gin.Context) {
	if rdb == nil {
		log.Panic("db cannot be nil")
	}

	return func(c *gin.Context) {
		email := c.GetString("email")
		_, err := rdb.Del(db_utils.Ctx, email).Result()

		if err != nil {
			log.Panic("Error deleting token from rdb")
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "ok"})
	}
}
