package model

type Background struct {
	ID             uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string `gorm:"unique;not null"`
	Description    string `gorm:"not null"`
	FeatID         uint
	Skill          []Skill `gorm:"many2many:skill_backgrounds;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AttributeBoost string  `gorm:"not null"`
}
