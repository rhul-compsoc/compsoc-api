package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// Gets all Members from members table.
//   - /member
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

// Gets a Member with id given in the parameter.
//   - /member/:member
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

// Posts a Member with data given in the body.
//   - /member
//
// body
//   - "id": int
//   - "student_id": string
//   - "student_email": string
//   - "first_name": string
//   - "last_name": string
//   - "active_member": bool
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

// Puts a Member with data given in the body.
//   - /member
//
// body
//   - "id": int
//   - "student_id": string
//   - "student_email": string
//   - "first_name": string
//   - "last_name": string
//   - "active_member": bool
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

// Patchers a Member with data given in the body.
//   - /member
//
// body
//   - "id": int
//   - "student_id": string
//   - "student_email": string
//   - "first_name": string
//   - "last_name": string
//   - "active_member": bool
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

// Deletes a Member with id given in the parameter.
//   - /member/:member
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
