package model

type CharacterItem struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterId uint `gorm:"not null;uniqueIndex:idx_character_item"`
	ItemId      uint `gorm:"not null;uniqueIndex:idx_character_item"`
	Quantity    uint `gorm:"not null;default=1"`

	Character Character `gorm:"foreignKey:CharacterId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Item      Item      `gorm:"foreignKey:ItemId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	ID          uint `json:"id" query:"id" form:"id"`
	CharacterID uint `json:"character_id" query:"character_id" form:"character_id"`
	ItemID      uint `json:"item_id" query:"item_id" form:"item_id"`
	Quantity    uint `json:"quantity" query:"quantity" form:"quantity" example:"1"`
}
