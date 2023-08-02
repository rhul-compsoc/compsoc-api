package database

import (
	"context"

	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
	"gorm.io/gorm"
)

const (
	MembersTable = "members"
	UsersTable   = "users"
	EventsTable  = "events"
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
	r := s.db.Table(MembersTable).Find(&m)

	return m, r.Error
}

// Get member, with their id, from members table.
func (s *Store) GetMember(id int) (models.MemberModel, error) {
	m := models.MemberModel{Id: id}
	r := s.db.Table(MembersTable).Find(&m).First(&m)

	return m, r.Error
}

// Add member to the members table.
func (s *Store) AddMember(m models.MemberModel) error {
	r := s.db.Table(MembersTable).Create(&m)

	return r.Error
}

// Update a member from the members table.
func (s *Store) UpdateMember(m models.MemberModel) error {
	r := s.db.Table(MembersTable).Save(&m)

	return r.Error
}

// Delete a member from the members table.
func (s *Store) DeleteMember(id int) error {
	m := models.MemberModel{Id: id}
	r := s.db.Table(MembersTable).Delete(&m)

	return r.Error
}

// Check a member exists in members table.
func (s *Store) CheckMember(id int) bool {
	m := models.MemberModel{Id: id}
	r := s.db.Table(MembersTable).Model(&m).First(&m)

	return r.Error == nil
}

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

// List persons from users table.
func (s *Store) ListEvent() ([]models.EventModel, error) {
	e := make([]models.EventModel, 0)
	r := s.db.Table(EventsTable).Find(&e)

	return e, r.Error
}

// Get event, with its id, from events table.
func (s *Store) GetEvent(id int) (models.EventModel, error) {
	e := models.EventModel{Id: id}
	r := s.db.Table(EventsTable).Find(&e).First(&e)

	return e, r.Error
}

// Add an event to the events table.
func (s *Store) AddEvent(p models.EventModel) error {
	r := s.db.Table(EventsTable).Create(&p)

	return r.Error
}

// Update an event from the events table.
func (s *Store) UpdateEvent(p models.EventModel) error {
	r := s.db.Table(EventsTable).Save(&p)

	return r.Error
}

// Delete an event from the events table.
func (s *Store) DeleteEvent(id int) error {
	e := models.EventModel{Id: id}
	r := s.db.Table(EventsTable).Delete(&e)

	return r.Error
}

// Check an event exists in events table.
func (s *Store) CheckEvent(id int) bool {
	e := models.EventModel{Id: id}
	r := s.db.Table(EventsTable).Model(&e).First(&e)

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
