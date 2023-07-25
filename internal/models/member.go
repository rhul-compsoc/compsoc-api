package models

type MemberModel struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	StudentId    string `gorm:"default:'student_id';type:varchar(50);not null"`
	StudentEmail string `gorm:"default:'student_email';type:varchar(50);not null"`
	FirstName    string `gorm:"default:'first_name';type:varchar(25);not null"`
	LastName     string `gorm:"default:'last_name';type:varchar(25);not null"`
	ActiveMember bool   `gorm:"default:false;type:boolean;not null"`
}

func (m *MemberModel) ToMemberPost() MemberPost {
	return MemberPost{
		Id:           m.Id,
		StudentId:    m.StudentId,
		StudentEmail: m.StudentEmail,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		ActiveMember: m.ActiveMember,
	}
}

type MemberPost struct {
	Id           int    `json:"id"`
	StudentId    string `json:"student_id"`
	StudentEmail string `json:"student_email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ActiveMember bool   `json:"active_member"`
}

func (m *MemberPost) ToMemberModel() MemberModel {
	return MemberModel{
		Id:           m.Id,
		StudentId:    m.StudentId,
		StudentEmail: m.StudentEmail,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		ActiveMember: m.ActiveMember,
	}
}
