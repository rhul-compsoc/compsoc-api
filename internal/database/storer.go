package database

import (
	"context"

	"gorm.io/gorm"
)

const (
	MembersTable  = "members"
	UsersTable    = "users"
	EventsTable   = "events"
	AdminsTable   = "admins"
	StudentsTable = "students"
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
