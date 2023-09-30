package database

import (
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// List admin tokens from admins table.
func (s *Store) ListAdmin() ([]models.Admin, error) {
	a := make([]models.Admin, 0)
	r := s.db.Table(AdminsTable).Find(&a)

	return a, r.Error
}

// Add an admin token to the admins table.
func (s *Store) AddAdmin(tok string) error {
	a := models.Admin{Token: tok}
	r := s.db.Table(AdminsTable).Create(&a)

	return r.Error
}

// Delete an admin token from the admins table.
func (s *Store) DeleteAdmin(tok string) error {
	a := models.Admin{Token: tok}
	r := s.db.Table(AdminsTable).Delete(&a)

	return r.Error
}

// Check an admin token exists in admins table.
func (s *Store) CheckAdmin(tok string) bool {
	a := models.Admin{Token: tok}
	r := s.db.Table(AdminsTable).Model(&a).First(&a)

	return r.Error == nil
}
