package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// Gets all Events from events table.
//   - /event
func EventList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := s.ListEvent()

		r := make([]models.EventPost, len(l))
		for i, event := range l {
			r[i] = event.ToEventPost()
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Gets an Event with id given in the parameter.
//   - /event/:event
func EventGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("event")
		id, err := strconv.Atoi(p)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckEvent(id)
		if !e {
			c.Status(http.StatusNotFound)
			return
		}

		g, err := s.GetEvent(id)
		r := g.ToEventPost()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// posts an Event with data given in the body.
//   - /event
//
// body
//   - "id": int
//   - "name": string
//   - "description": string
//   - "date": Time
//   - "time": Date
//   - "attendance": int
//   - "members_only": bool
func EventPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.DefaultEventPost()
		c.Bind(&b)
		m := b.ToEventModel()

		if m.Id != 0 {
			e := s.CheckEvent(m.Id)
			if e {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		err := s.AddEvent(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Puts an Event with data given in the body.
//   - /event
//
// body
//   - "id": int
//   - "name": string
//   - "description": string
//   - "date": Time
//   - "time": Date
//   - "attendance": int
//   - "members_only": bool
func EventPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.DefaultEventPost()
		c.Bind(&b)
		m := b.ToEventModel()

		e := s.CheckEvent(m.Id)
		if !e {
			err := s.AddEvent(m)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			c.Status(http.StatusOK)
			return
		}

		err := s.UpdateEvent(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Patchers an Event with data given in the body.
//   - /event
//
// body
//   - "id": int
//   - "name": string
//   - "description": string
//   - "date": Time
//   - "time": Date
//   - "attendance": int
//   - "members_only": bool
func EventPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := models.DefaultEventPost()
		c.Bind(&b)
		m := b.ToEventModel()

		e := s.CheckEvent(m.Id)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdateEvent(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// Deletes an Event with id given in the parameter.
//   - /event/:event
func EventDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("event")
		id, err := strconv.Atoi(p)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckEvent(id)
		if !e {
			c.Status(http.StatusNotFound)
		}

		err = s.DeleteEvent(id)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}
