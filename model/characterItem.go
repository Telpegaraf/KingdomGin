package model

type CharacterItem struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterId uint `gorm:"not null;uniqueIndex:idx_character_item"`
	ItemId      uint `gorm:"not null;uniqueIndex:idx_character_item"`
	Quantity    uint `gorm:"not null;default=1"`

	Character Character `gorm:"foreignKey:CharacterId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Item      Item      `gorm:"foreignKey:ItemId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
