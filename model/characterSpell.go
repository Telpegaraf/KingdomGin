package model

type CharacterSpell struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterID uint `gorm:"foreignKey:CharacterID;references:ID"`
	SpellID     uint `gorm:"foreignKey:SpellID;references:ID"`
}

type CharacterSpellCreate struct {
	CharacterID uint `json:"character_id" query:"character_id" form:"character_id" binding:"required"`
	SpellID     uint `json:"spell_id" query:"spell_id" form:"spell_id" binding:"required"`
}

type CharacterSpellUpdate struct {
	SpellID uint `json:"spell_id" query:"spell_id" form:"spell_id"`
}

type CharacterSpellExternal struct {
	ID      uint `json:"id" query:"id" form:"id"`
	SpellID uint `json:"spell_id" query:"spell_id" form:"spell_id"`
}
