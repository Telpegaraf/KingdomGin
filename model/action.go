package model

type Action struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(127);not null"`

	Spells []Spell `gorm:"many2many:spell_actions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
