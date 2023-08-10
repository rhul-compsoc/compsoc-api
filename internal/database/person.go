package database

import "github.com/rhul-compsoc/compsoc-api-go/internal/models"

// List persons from users table.
func (s *Store) ListPerson() ([]models.PersonModel, error) {
	p := make([]models.PersonModel, 0)
	r := s.db.Table(UsersTable).Find(&p)

	return p, r.Error
}

// Get person, with their id, from users table.
func (s *Store) GetPerson(id int) (models.PersonModel, error) {
	p := models.PersonModel{Id: id}
	r := s.db.Table(UsersTable).Find(&p).First(&p)

	return p, r.Error
}

// Add person to the users table
func (s *Store) AddPerson(p models.PersonModel) error {
	r := s.db.Table(UsersTable).Create(&p)

	return r.Error
}

// Update a person from the users table.
func (s *Store) UpdatePerson(p models.PersonModel) error {
	r := s.db.Table(UsersTable).Save(&p)

	return r.Error
}

// Delete a person from the users table.
func (s *Store) DeletePerson(id int) error {
	p := models.PersonModel{Id: id}
	r := s.db.Table(UsersTable).Delete(&p)

	return r.Error
}

// Check a person exists in users table.
func (s *Store) CheckPerson(id int) bool {
	p := models.PersonModel{Id: id}
	r := s.db.Table(UsersTable).Model(&p).First(&p)

	return r.Error == nil
}
