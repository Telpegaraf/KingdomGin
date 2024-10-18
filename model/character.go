package model

type Character struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"not null;type:varchar(120)"`
	Alias            string `gorm:"type:varchar(120)"`
	LastName         string `gorm:"type:varchar(120)" json:"last_name"`
	Level            int8   `gorm:"default:1"`
	UserID           uint
	RaceID           uint
	AncestryID       uint
	BackgroundID     uint
	CharacterClassID uint
	Attribute        Attribute        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterSpell   []CharacterSpell `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterItem    []CharacterItem  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Slot             []Slot           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterBoost   CharacterBoost   `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	CharacterDefence CharacterDefence `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	CharacterFeat    CharacterFeat    `gorm:"constraint:OnUpdate:CASCADE;onDelete:CASCADE;"`
}

type CreateCharacter struct {
	Name             string `binding:"required" json:"name" query:"name" form:"name"`
	Alias            string `json:"alias" query:"alias" form:"alias"`
	LastName         string `json:"last_name" query:"last_name" form:"last_name"`
	CharacterClassID uint   `json:"character_class_id" query:"character_class_id" form:"character_class_id"`
	RaceID           uint   `json:"race_id" query:"race_id" form:"race_id"`
	AncestryID       uint   `json:"ancestry_id" query:"ancestry_id" form:"ancestry_id" binding:"required"`
	BackgroundID     uint   `json:"background_id" query:"background_id" form:"background_id" binding:"required"`
}

type CharacterUpdate struct {
	Name     string `json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string `json:"last_name" query:"last_name" form:"last_name"`
	Level    int8   `json:"level" query:"level" form:"level"`
}

type CharacterExternal struct {
	ID               uint            `json:"id"`
	Name             string          `binding:"required" json:"name" query:"name" form:"name"`
	Alias            string          `json:"alias" query:"alias" form:"alias"`
	LastName         string          `json:"last_name" query:"last_name" form:"last_name"`
	Level            int8            `json:"level" query:"level" form:"level"`
	RaceID           uint            `json:"race_id" query:"race_id" form:"race_id"`
	UserID           uint            `json:"user_id" query:"user_id" form:"user_id"`
	Attribute        Attribute       `json:"attribute" query:"attribute" form:"attribute"`
	CharacterItem    []CharacterItem `json:"character_item" query:"character_item" form:"character_item"`
	Slot             []Slot          `json:"slot" query:"slot" form:"slot"`
	CharacterBoost   CharacterBoost  `json:"character_boost" query:"character_boost" form:"character_boost"`
	CharacterClassID uint            `json:"character_class_id" query:"character_class_id" form:"character_class_id"`
	AncestryID       uint            `json:"ancestry_id" query:"ancestry_id" form:"ancestry_id"`
	BackgroundID     uint            `json:"background_id" query:"background_id" form:"background_id"`
}
