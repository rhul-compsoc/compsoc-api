package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// Gets all Persons from users table.
//   - /user
func UserList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := s.ListUser()

		r := make([]models.UserPost, len(l))
		for i, user := range l {
			r[i] = user.ToPost()
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Gets a Person with id given in the parameter.
//   - /user/:user
func UserGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("user")
		id, err := strconv.Atoi(p)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckPerson(id)
		if !e {
			c.Status(http.StatusNotFound)
			return
		}

		g, err := s.GetUser(id)
		r := g.ToPost()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Posts a Member with data given in the body.
//   - /user
//
// body
//   - "id": int
//   - "name": string
func UserPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b models.UserPost
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckPerson(m.Id)
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.AddUser(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Puts a Member with data given in the body.
//   - /user
//
// body
//   - "id": int
//   - "name": string
func UserPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b models.UserPost
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckPerson(m.Id)
		if !e {
			err := s.AddUser(m)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			c.Status(http.StatusOK)
			return
		}

		err := s.UpdateUser(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Patchs a Member with data given in the body.
//   - /user
//
// body
//   - "id": int
//   - "name": string
func UserPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b models.UserPost
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckPerson(m.Id)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdateUser(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Deletes a Person with id given in the parameter.
//   - /user/:user
func UserDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("user")
		id, err := strconv.Atoi(p)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckPerson(id)
		if !e {
			c.Status(http.StatusNotFound)
		}

		err = s.DeleteUser(id)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}
