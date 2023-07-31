package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

func PersonList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := s.ListPerson()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

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

		r, err := s.GetPerson(id)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

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
