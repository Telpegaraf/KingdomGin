package model

type Spell struct {
	ID             uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string
	Description    string `gorm:"type:text"`
	Component      string `gorm:"type:text"`
	Range          string `gorm:"type:varchar(255)"`
	Area           string `gorm:"type:varchar(255)"`
	Duration       string `gorm:"type:varchar(255)"`
	Target         string `gorm:"type:varchar(255)"`
	Source         string `gorm:"type:varchar(255)"`
	Rank           uint8
	Ritual         bool             `gorm:"type:bool;default:false"`
	School         School           `gorm:"type:school"`
	CharacterSpell []CharacterSpell `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cast           string
	Traditional    []Tradition `gorm:"many2many:spell_traditions"`
	Traits         []Trait     `gorm:"many2many:spell_traits;"`
}
