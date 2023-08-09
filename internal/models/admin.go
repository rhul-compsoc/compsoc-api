package models

type Admin struct {
	Token string `gorm:"primaryKey;default:'token';type:varchar(64);not null"`
}
