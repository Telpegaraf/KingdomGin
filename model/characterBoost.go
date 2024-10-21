package model

type CharacterBoost struct {
	ID              uint  `gorm:"primary_key;AUTO_INCREMENT"`
	AncestryBoost   uint8 `gorm:"default:2"`
	BackgroundBoost bool  `gorm:"default:true"`
	ClassBoost      bool  `gorm:"default:true"`
	FreeBoost       uint8 `gorm:"default:1"`
	CharacterID     uint  `gorm:"foreignKey:CharacterID;unique"`
}

type CreateCharacterBoost struct {
	CharacterID uint
}

type UpdateCharacterBoost struct {
}

type CharacterBoostExternal struct {
	ID          uint `json:"id" query:"id" form:"id"`
	CharacterID uint
}
