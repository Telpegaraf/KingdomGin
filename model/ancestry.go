package model

type Ancestry struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `gorm:"unique"`
	Description string `gorm:"type:text"`
	RaceID      uint
}
