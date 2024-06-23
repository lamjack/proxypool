package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Health(startedAt time.Time) func(c *gin.Context) {
	return func(c *gin.Context) {
		now := time.Now().UTC()
		uptime := now.Sub(startedAt)
		c.JSON(
			http.StatusOK, gin.H{
				"startedAt": startedAt.String(),
				"uptime":    uptime.String(),
				"status":    "OK",
				"ipAddress": c.ClientIP(),
			},
		)
	}
}
