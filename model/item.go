package model

type Item struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique;type:varchar(127)"`
	Description string `gorm:"type:text"`
}
