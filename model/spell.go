package model

type Spell struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string
	Description string
	Tradition   Tradition `gorm:"type:tradition"`
	School      School    `gorm:"type:school"`
}
