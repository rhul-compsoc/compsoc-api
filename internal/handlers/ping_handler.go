package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
)

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
