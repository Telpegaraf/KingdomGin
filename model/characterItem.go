package model

type CharacterItem struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterID  uint   `gorm:"not null;uniqueIndex:idx_character_item"`
	ItemID       uint   `gorm:"not null;uniqueIndex:idx_character_item"`
	Quantity     uint   `gorm:"not null;default=1"`
	Armor        []Slot `gorm:"foreignKey:ArmorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FirstWeapon  []Slot `gorm:"foreignKey:FirstWeaponID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SecondWeapon []Slot `gorm:"foreignKey:SecondWeaponID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Character Character `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Item      Item      `gorm:"foreignKey:ItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CreateCharacterItem struct {
	CharacterID uint `json:"character_id" query:"character_id" binding:"required" form:"character_id"`
	ItemID      uint `json:"item_id" query:"item_id" binding:"required" form:"item_id"`
	Quantity    uint `json:"quantity" query:"quantity" form:"quantity" example:"1"`
}

type UpdateCharacterItem struct {
	CharacterID uint `json:"character_id" query:"character_id" form:"character_id"`
	ItemID      uint `json:"item_id" query:"item_id" form:"item_id"`
	Quantity    uint `json:"quantity" query:"quantity" binding:"required" form:"quantity" example:"1"`
}

type CharacterItemExternal struct {
	ID            uint    `json:"id" query:"id" form:"id"`
	CharacterID   uint    `json:"character_id" query:"character_id" form:"character_id"`
	CharacterName string  `json:"character_name" query:"character_name" form:"character_name"`
	Quantity      uint    `json:"quantity" query:"quantity" form:"quantity" example:"1"`
	ItemID        uint    `json:"itemID" query:"item_id" form:"item_id"`
	ItemName      string  `json:"item_name" query:"item_name" form:"item_name"`
	Bulk          float64 `json:"bulk" query:"bulk" form:"bulk"`
}
