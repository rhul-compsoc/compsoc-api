package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
)

// Pings the API.
//   - /ping
//
// "pong" will be in response if database connection is ok.
func PingGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := "pong"

		err := s.Ping()
		if err != nil {
			resp = "boom"
		}

		c.JSON(http.StatusOK, gin.H{
			"ping": resp,
		})
	}
}

// Pings the API.
//   - /ping/admin
//
// "admin" will be in response if the user is authorized.
func AdminPingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := "admin"

		c.JSON(http.StatusOK, gin.H{
			"ping": resp,
		})
	}
}
