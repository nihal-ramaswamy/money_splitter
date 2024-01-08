package middleware

import (
	"fmt"
	"money_splitter/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secret := utils.GetDotEnvVariable("SECRET_KEY")
	signingKey := []byte(secret)

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
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

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			c.Set("email", claims["email"])
			c.Set("authenticated", true)
		}

		c.Next()
	}
}
