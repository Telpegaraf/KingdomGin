package model

type Attribute struct {
	ID           uint  `gorm:"primary_key;AUTO_INCREMENT"`
	Strength     uint8 `gorm:"default:10;not null;"`
	Dexterity    uint8 `gorm:"default:10;not null;"`
	Constitution uint8 `gorm:"default:10;not null"`
	Intelligence uint8 `gorm:"default:10;not null"`
	Wisdom       uint8 `gorm:"default:10;not null"`
	Charisma     uint8 `gorm:"default:10;not null"`
	CharacterID  uint  `gorm:"unique"`
}

type CreateAttribute struct {
	Strength     uint8 `json:"strength" query:"strength" form:"strength"`
	Dexterity    uint8 `json:"dexterity" query:"dexterity" form:"dexterity"`
	Constitution uint8 `json:"constitution" query:"constitution" form:"constitution"`
	Intelligence uint8 `json:"intelligence" query:"intelligence" form:"intelligence"`
	Wisdom       uint8 `json:"wisdom" query:"wisdom" form:"wisdom"`
	Charisma     uint8 `json:"charisma" query:"charisma" form:"charisma"`
	CharacterID  uint
}

type UpdateAttribute struct {
	Strength     *uint8 `json:"strength" query:"strength" form:"strength"`
	Dexterity    *uint8 `json:"dexterity" query:"dexterity" form:"dexterity"`
	Constitution *uint8 `json:"constitution" query:"constitution" form:"constitution"`
	Intelligence *uint8 `json:"intelligence" query:"intelligence" form:"intelligence"`
	Wisdom       *uint8 `json:"wisdom" query:"wisdom" form:"wisdom"`
	Charisma     *uint8 `json:"charisma" query:"charisma" form:"charisma"`
	CharacterID  uint
}

type AttributeExternal struct {
	ID           uint  `json:"id" query:"id" form:"id"`
	Strength     uint8 `json:"strength" query:"strength" form:"strength"`
	Dexterity    uint8 `json:"dexterity" query:"dexterity" form:"dexterity"`
	Constitution uint8 `json:"constitution" query:"constitution" form:"constitution"`
	Intelligence uint8 `json:"intelligence" query:"intelligence" form:"intelligence"`
	Wisdom       uint8 `json:"wisdom" query:"wisdom" form:"wisdom"`
	Charisma     uint8 `json:"charisma" query:"charisma" form:"charisma"`
	CharacterID  uint
}
