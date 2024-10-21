package model

type CharacterBoost struct {
	ID              uint  `gorm:"primary_key;AUTO_INCREMENT"`
	AncestryBoost   bool  `gorm:"default:true"`
	BackgroundBoost bool  `gorm:"default:0"`
	FreeBoost       uint8 `gorm:"default:0"`
	CharacterID     uint  `gorm:"unique"`
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
