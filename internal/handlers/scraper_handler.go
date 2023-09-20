package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScraperPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var b interface{}
		err := c.ShouldBindJSON(&b)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		fmt.Println(b)
		c.Status(http.StatusOK)
	}
}
