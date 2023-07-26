package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
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
			return
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
		b := models.MemberPost{}
		c.Bind(&b)
		m := b.ToMemberModel()

		e := s.CheckMember(m.Id)
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.AddMember(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

func MemberPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.MemberPost{}
		c.Bind(&b)
		m := b.ToMemberModel()

		e := s.CheckMember(m.Id)
		if !e {
			err := s.AddMember(m)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			c.Status(http.StatusOK)
			return
		}

		err := s.UpdateMember(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

func MemberPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.MemberPost{}
		c.Bind(&b)
		m := b.ToMemberModel()

		e := s.CheckMember(m.Id)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdateMember(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

func MemberDelete(s *database.Store) gin.HandlerFunc {
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

		err = s.DeleteMember(id)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}
