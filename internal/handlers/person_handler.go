package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// Gets all Persons from users table.
//   - /person
func PersonList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := s.ListPerson()

		r := make([]models.PersonPost, len(l))
		for i, person := range l {
			r[i] = person.ToPersonPost()
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Gets a Person with id given in the parameter.
//   - /person/:person
func PersonGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("person")
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

		g, err := s.GetPerson(id)
		r := g.ToPersonPost()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Posts a Member with data given in the body.
//   - /person
//
// body
//   - "id": int
//   - "name": string
func PersonPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.PersonPost{}
		c.Bind(&b)
		p := b.ToPersonModel()

		e := s.CheckPerson(p.Id)
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.AddPerson(p)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Puts a Member with data given in the body.
//   - /person
//
// body
//   - "id": int
//   - "name": string
func PersonPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.PersonPost{}
		c.Bind(&b)
		p := b.ToPersonModel()

		e := s.CheckPerson(p.Id)
		if !e {
			err := s.AddPerson(p)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			c.Status(http.StatusOK)
			return
		}

		err := s.UpdatePerson(p)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Patchs a Member with data given in the body.
//   - /person
//
// body
//   - "id": int
//   - "name": string
func PersonPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.PersonPost{}
		c.Bind(&b)
		m := b.ToPersonModel()

		e := s.CheckPerson(m.Id)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdatePerson(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Deletes a Person with id given in the parameter.
//   - /person/:person
func PersonDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("person")
		id, err := strconv.Atoi(p)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckPerson(id)
		if !e {
			c.Status(http.StatusNotFound)
		}

		err = s.DeletePerson(id)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}
