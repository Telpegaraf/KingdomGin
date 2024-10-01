package model

type Domain struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique;not null;type:varchar(120)"`
	Description string `gorm:"type:text"`
}
