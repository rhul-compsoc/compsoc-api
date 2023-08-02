package models

import "github.com/rhul-compsoc/compsoc-api-go/pkg/util"

type EventModel struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"default:'event_name';type:varchar(50);not null"`
	Description string `gorm:"default:'event_description';type:text;not null"`
	Date        string `gorm:"default:'00/00/0000';type:varchar(10);not null"`
	Time        string `gorm:"default:'00:00';type:varchar(5);not null"`
	Attendance  int    `gorm:"default:0;type:integer;not null"`
	MembersOnly bool   `gorm:"default:false;type:boolean;not null"`
}

func (e *EventModel) ToEventPost() EventPost {
	d, err := NewDate(e.Date)
	util.ErrLog(err)

	t, err := NewTime(e.Time)
	util.ErrLog(err)

	return EventPost{
		Id:          e.Id,
		Name:        e.Name,
		Description: e.Description,
		Date:        d,
		Time:        t,
		Attendance:  e.Attendance,
		MembersOnly: e.MembersOnly,
	}
}

func DefaultEventPost() EventPost {
	d, err := NewDate("00/00/0000")
	util.LogErr(err)

	t, err := NewTime("00:00")
	util.LogErr(err)

	return EventPost{
		Name:        "name",
		Description: "description",
		Date:        d,
		Time:        t,
		Attendance:  0,
		MembersOnly: false,
	}
}

type EventPost struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        Date   `json:"date"`
	Time        Time   `json:"time"`
	Attendance  int    `json:"attendance"`
	MembersOnly bool   `json:"members_only"`
}

func (e *EventPost) ToEventModel() EventModel {
	return EventModel{
		Id:          e.Id,
		Name:        e.Name,
		Description: e.Description,
		Date:        e.Date.String(),
		Time:        e.Time.String(),
		Attendance:  e.Attendance,
		MembersOnly: e.MembersOnly,
	}
}
