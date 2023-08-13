package models

import "github.com/rhul-compsoc/compsoc-api-go/pkg/util"

type StudentPost struct {
	Id           int
	Email        string
	Verified     bool
	DateVerified Date
}

func (s *StudentPost) ToModel() StudentModel {
	d := s.DateVerified.String()

	return StudentModel{
		Id:           s.Id,
		Email:        s.Email,
		Verified:     s.Verified,
		DateVerified: d,
	}
}

type StudentModel struct {
	Id           int
	Email        string
	Verified     bool
	DateVerified string
}

func (s *StudentModel) ToPost() StudentPost {
	d, err := NewDate(s.DateVerified)
	util.ErrLog(err)

	return StudentPost{
		Id:           s.Id,
		Email:        s.Email,
		Verified:     s.Verified,
		DateVerified: d,
	}
}
