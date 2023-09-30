package decorators

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
)

func AdminAuth() DecoratorFunc {
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

func CommitteeAuth(s *database.Store) DecoratorFunc {
	return func(fn gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			h := c.Request.Header.Get("Authorization")
			a := strings.Fields(h)

			if a[0] != "Admin" {
				c.Status(http.StatusUnauthorized)
				return
			}

			e := s.CheckAdmin(a[1])
			if !e {
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
