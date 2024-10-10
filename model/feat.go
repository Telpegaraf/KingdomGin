package model

type Feat struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `gorm:"unique;not null"`
	Description string
	Background  []Background
}
