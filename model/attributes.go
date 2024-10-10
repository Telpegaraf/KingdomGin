package model

type Attributes struct {
	ID           uint `gorm:"primary_key;AUTO_INCREMENT"`
	Strength     int8 `gorm:"default:10"`
	Dexterity    int8 `gorm:"default:10"`
	Constitution int8 `gorm:"default:10"`
	Intelligence int8 `gorm:"default:10"`
	Wisdom       int8 `gorm:"default:10"`
	Charisma     int8 `gorm:"default:10"`
	CharacterID  uint `gorm:"unique"`
}
