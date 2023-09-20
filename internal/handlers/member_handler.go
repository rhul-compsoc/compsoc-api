package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// Gets all Members from members table.
//   - /api/v1/member
func MemberList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := s.ListMember()

		r := make([]models.MemberPost, len(l))
		for i, member := range l {
			r[i] = member.ToPost()
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Gets a Member with id given in the parameter.
//   - /api/v1/member/:member
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

		g, err := s.GetMember(id)
		r := g.ToPost()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Posts a Member with data given in the body.
//   - /api/v1/member
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
		var b models.MemberPost
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckPerson(m.Id)
		if !e {
			p := models.UserModel{
				Id:   m.Id,
				Name: m.FirstName,
			}
			err := s.AddUser(p)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		e = s.CheckMember(m.Id)
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
//   - /api/v1/member
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
		var b models.MemberPost
		c.ShouldBindJSON(&b)
		m := b.ToModel()

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
//   - /api/v1/member
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
		var b models.MemberPost
		c.ShouldBindJSON(&b)
		m := b.ToModel()

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
//   - /api/v1/member/:member
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
