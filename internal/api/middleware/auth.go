package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"money_splitter/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"

	db_utils "money_splitter/internal/db"
)

func AuthMiddleware(pdb *sql.DB, rdb *redis.Client) gin.HandlerFunc {
	secret := utils.GetDotEnvVariable("SECRET_KEY")
	signingKey := []byte(secret)

	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No token found"})
			return
		}

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error in parsing token")
			}
			return signingKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		email := parsedToken.Claims.(jwt.MapClaims)["email"].(string)

		_, err = rdb.Get(db_utils.Ctx, email).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		user, err := db_utils.GetUserFromEmail(pdb, email)

		if err != nil {
			log.Panic(err)
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			c.Set("email", claims["email"])
			c.Set("userID", user.ID)
			c.Set("authenticated", true)
		}

		c.Next()
	}
}
