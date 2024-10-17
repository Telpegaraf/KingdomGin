package model

type CharacterBoost struct {
	ID           uint  `gorm:"primary_key;AUTO_INCREMENT"`
	Strength     uint8 `gorm:"default:0;not null;"`
	Dexterity    uint8 `gorm:"default:0;not null;"`
	Constitution uint8 `gorm:"default:0;not null"`
	Intelligence uint8 `gorm:"default:0;not null"`
	Wisdom       uint8 `gorm:"default:0;not null"`
	Charisma     uint8 `gorm:"default:0;not null"`
	Free         uint8 `gorm:"default:0;not null"`
	CharacterID  uint  `gorm:"unique"`
}

type CreateCharacterBoost struct {
	Strength     *uint8 `json:"strength" query:"strength" form:"strength"`
	Dexterity    *uint8 `json:"dexterity" query:"dexterity" form:"dexterity"`
	Constitution *uint8 `json:"constitution" query:"constitution" form:"constitution"`
	Intelligence *uint8 `json:"intelligence" query:"intelligence" form:"intelligence"`
	Wisdom       *uint8 `json:"wisdom" query:"wisdom" form:"wisdom"`
	Charisma     *uint8 `json:"charisma" query:"charisma" form:"charisma"`
	Free         *uint8 `json:"free" query:"free" form:"free"`
	CharacterID  uint
}

type UpdateCharacterBoost struct {
	Strength     *uint8 `json:"strength" query:"strength" form:"strength" example:"10"`
	Dexterity    *uint8 `json:"dexterity" query:"dexterity" form:"dexterity" example:"10"`
	Constitution *uint8 `json:"constitution" query:"constitution" form:"constitution" example:"10"`
	Intelligence *uint8 `json:"intelligence" query:"intelligence" form:"intelligence" example:"10"`
	Wisdom       *uint8 `json:"wisdom" query:"wisdom" form:"wisdom" example:"10"`
	Charisma     *uint8 `json:"charisma" query:"charisma" form:"charisma" example:"10"`
	Free         *uint8 `json:"free" query:"free" form:"free"`
}

type CharacterBoostExternal struct {
	ID           uint  `json:"id" query:"id" form:"id"`
	Strength     uint8 `json:"strength" query:"strength" form:"strength"`
	Dexterity    uint8 `json:"dexterity" query:"dexterity" form:"dexterity"`
	Constitution uint8 `json:"constitution" query:"constitution" form:"constitution"`
	Intelligence uint8 `json:"intelligence" query:"intelligence" form:"intelligence"`
	Wisdom       uint8 `json:"wisdom" query:"wisdom" form:"wisdom"`
	Charisma     uint8 `json:"charisma" query:"charisma" form:"charisma"`
	Free         uint8 `json:"free" query:"free" form:"free"`
	CharacterID  uint
}
