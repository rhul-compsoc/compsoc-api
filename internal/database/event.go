package database

import "github.com/rhul-compsoc/compsoc-api-go/internal/models"

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
func (s *Store) AddEvent(e models.EventModel) error {
	r := s.db.Table(EventsTable).Create(&e)

	return r.Error
}

// Update an event from the events table.
func (s *Store) UpdateEvent(e models.EventModel) error {
	r := s.db.Table(EventsTable).Save(&e)

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
