package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
)

// Gets all Events from events table.
//   - /event
func EventList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Gets an Event with id given in the parameter.
//   - /event/:event
func EventGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

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

	}
}

// Deletes an Event with id given in the parameter.
//   - /event/:event
func EventDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
