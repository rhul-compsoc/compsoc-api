package models

type UserModel struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"default:'name';type:varchar(50);not null"`
}

func (p *UserModel) ToPost() UserPost {
	return UserPost{
		Id:   p.Id,
		Name: p.Name,
	}
}

type UserPost struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (p *UserPost) ToModel() UserModel {
	return UserModel{
		Id:   p.Id,
		Name: p.Name,
	}
}
