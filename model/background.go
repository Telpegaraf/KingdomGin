package model

type Background struct {
	ID             uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string `gorm:"unique;not null"`
	Description    string `gorm:"not null"`
	FeatID         uint
	FirstSkillID   uint
	SecondSkillID  uint
	AttributeBoost string `gorm:"not null"`
}
