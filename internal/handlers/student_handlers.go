package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
)

// Gets all Students from members table.
//   - /student
func StudentList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Gets a Student with id given in the parameter.
//   - /student/:student
func StudentGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Posts a Student with data given in the body.
//   - /student
//
// body
//   - "id": int
func StudentPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Puts a Student with data given in the body.
//   - /student
//
// body
//   - "id": int
func StudentPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Patchers a Student with data given in the body.
//   - /student
//
// body
//   - "id": int
func StudentPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Deletes a Student with id given in the parameter.
//   - /student/:student
func StudentDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
