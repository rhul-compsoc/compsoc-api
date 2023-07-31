package models

type PersonModel struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"default:'name';type:varchar(50);not null"`
}

func (p *PersonModel) ToPersonPost() PersonPost {
	return PersonPost{
		Id:   p.Id,
		Name: p.Name,
	}
}

type PersonPost struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (p *PersonPost) ToPersonModel() PersonModel {
	return PersonModel{
		Id:   p.Id,
		Name: p.Name,
	}
}
