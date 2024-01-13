package group_api

import (
	"database/sql"
	"log"
	"money_splitter/internal/db"
	"money_splitter/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewGroupHandler(pdb *sql.DB, rdb *redis.Client) gin.HandlerFunc {
	if pdb == nil || rdb == nil {
		log.Panic("db cannot be null")
	}

	return func(c *gin.Context) {
		userId := c.GetString("userID")

		var group dto.Group
		if err := c.ShouldBindJSON(&group); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
		}

		group = group.SetGroupAdmin(userId)

		groupId := db.RegisterNewGroup(pdb, group)
		c.JSON(http.StatusAccepted, gin.H{"id": groupId})
	}
}
