package models

import "github.com/rhul-compsoc/compsoc-api-go/pkg/util"

type EventModel struct {
	Id          int
	Name        string
	Description string
	Date        string
	Time        string
	Attendance  int
	MembersOnly bool
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
