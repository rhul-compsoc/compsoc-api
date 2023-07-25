package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
)

func MemberList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := s.ListMember()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, m)
	}
}

func MemberGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("member")
		id, err := strconv.Atoi(p)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckMember(id)
		if !e {
			c.Status(http.StatusNotFound)
		}

		m, err := s.GetMember(id)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, m)
	}
}

func MemberPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func MemberPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func MemberPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func MemberDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
