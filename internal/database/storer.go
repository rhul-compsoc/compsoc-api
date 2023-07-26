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
