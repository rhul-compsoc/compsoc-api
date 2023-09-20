package decorators

import (
	"fmt"
	"net/http"
	"os"

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

func ScraperAuth() DecoratorFunc {
	return func(fn gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			a := c.Request.Header.Get("X-Auth-Token")
			fmt.Println(a)
			if a != os.Getenv("SCRAPER_TOK") {
				c.Status(http.StatusUnauthorized)
				return
			}

			fn(c)
		}
	}
}
