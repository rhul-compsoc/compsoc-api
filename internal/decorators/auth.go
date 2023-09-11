package decorators

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() DecoratorFunc {
	return func(fn gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			a := c.Request.Header.Get("auth")
			if a != "admin" {
				c.Status(http.StatusUnauthorized)
				return
			}

			fn(c)
		}
	}
}
