package database

import "github.com/rhul-compsoc/compsoc-api-go/internal/models"

// List persons from users table.
func (s *Store) ListUser() ([]models.UserModel, error) {
	m := make([]models.UserModel, 0)
	r := s.db.Table(UsersTable).Find(&m)

	return m, r.Error
}

// Get person, with their id, from users table.
func (s *Store) GetUser(id int) (models.UserModel, error) {
	p := models.UserModel{Id: id}
	r := s.db.Table(UsersTable).Find(&p).First(&p)

	return p, r.Error
}

// Add person to the users table
func (s *Store) AddUser(m models.UserModel) error {
	r := s.db.Table(UsersTable).Create(&m)

	return r.Error
}

// Update a person from the users table.
func (s *Store) UpdateUser(m models.UserModel) error {
	r := s.db.Table(UsersTable).Save(&m)

	return r.Error
}

// Delete a person from the users table.
func (s *Store) DeleteUser(id int) error {
	p := models.UserModel{Id: id}
	r := s.db.Table(UsersTable).Delete(&p)

	return r.Error
}

// Check a person exists in users table.
func (s *Store) CheckPerson(id int) bool {
	p := models.UserModel{Id: id}
	r := s.db.Table(UsersTable).Model(&p).First(&p)

	return r.Error == nil
}
