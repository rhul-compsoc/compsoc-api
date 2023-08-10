package database

import "github.com/rhul-compsoc/compsoc-api-go/internal/models"

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
