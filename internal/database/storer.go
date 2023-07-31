package database

import (
	"context"

	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
	"gorm.io/gorm"
)

// Stores a *gorm.DB and mutex.
type Store struct {
	db *gorm.DB
}

// Create a new pointer to a Store.
func New() *Store {
	return &Store{
		db: makeDb(),
	}
}

// List members from members table.
func (s *Store) ListMember() ([]models.MemberModel, error) {
	m := make([]models.MemberModel, 0)
	r := s.db.Table("members").Find(&m)

	return m, r.Error
}

// Get member, with their id, from members table.
func (s *Store) GetMember(id int) (models.MemberModel, error) {
	m := models.MemberModel{Id: id}
	r := s.db.Table("members").Find(&m).First(&m)

	return m, r.Error
}

func (s *Store) AddMember(m models.MemberModel) error {
	r := s.db.Table("members").Create(&m)

	return r.Error
}

// Update a member from the members table.
func (s *Store) UpdateMember(m models.MemberModel) error {
	r := s.db.Table("members").Save(&m)

	return r.Error
}

// Delete a member from the members table.
func (s *Store) DeleteMember(id int) error {
	m := models.MemberModel{Id: id}
	r := s.db.Table("members").Delete(&m)

	return r.Error
}

// Check a member exists in members table.
func (s *Store) CheckMember(id int) bool {
	m := models.MemberModel{Id: id}
	r := s.db.Table("members").Model(&m).First(&m)

	return r.Error == nil
}

// List persons from users table.
func (s *Store) ListPerson() ([]models.PersonModel, error) {
	p := make([]models.PersonModel, 0)
	r := s.db.Table("users").Find(&p)

	return p, r.Error
}

// Get person, with their id, from users table.
func (s *Store) GetPerson(id int) (models.PersonModel, error) {
	p := models.PersonModel{Id: id}
	r := s.db.Table("users").Find(&p).First(&p)

	return p, r.Error
}

func (s *Store) AddPerson(p models.PersonModel) error {
	r := s.db.Table("users").Create(&p)

	return r.Error
}

// Update a person from the users table.
func (s *Store) UpdatePerson(p models.PersonModel) error {
	r := s.db.Table("users").Save(&p)

	return r.Error
}

// Delete a person from the users table.
func (s *Store) DeletePerson(id int) error {
	m := models.PersonModel{Id: id}
	r := s.db.Table("users").Delete(&m)

	return r.Error
}

// Check a person exists in users table.
func (s *Store) CheckPerson(id int) bool {
	m := models.PersonModel{Id: id}
	r := s.db.Table("users").Model(&m).First(&m)

	return r.Error == nil
}

// Ping the database.
func (s *Store) Ping() error {
	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}
