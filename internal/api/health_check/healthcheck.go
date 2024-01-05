package healthcheck_api

import "github.com/gin-gonic/gin"

func HealthCheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	}
}
