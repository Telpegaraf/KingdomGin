package model

type SpellCharacterClass struct {
	ID               uint  `gorm:"primary_key;AUTO_INCREMENT"`
	level            uint8 `gorm:"default:1"`
	spellCount       uint8 `gorm:"default:1"`
	CharacterClassID *uint
	CharacterClass   CharacterClass `gorm:"foreignKey:CharacterClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
