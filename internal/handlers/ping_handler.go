package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	}
}
