package model

type Spell struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string
	Description string      `gorm:"type:text"`
	Range       string      `gorm:"type:varchar(127)"`
	Area        string      `gorm:"type:varchar(127)"`
	Duration    string      `gorm:"type:varchar(127)"`
	School      School      `gorm:"type:school"`
	Traditional []Tradition `gorm:"many2many:spell_traditions"`
	Traits      []Trait     `gorm:"many2many:spell_traits;"`
}
