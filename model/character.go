package model

type Character struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"not null;type:varchar(120)"`
	Alias            string `gorm:"type:varchar(120)"`
	LastName         string `gorm:"type:varchar(120)" json:"last_name"`
	Level            int8   `gorm:"default:1"`
	UserID           uint
	Attribute        Attribute        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterItem    []CharacterItem  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Slot             []Slot           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterBoost   CharacterBoost   `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	CharacterDefence CharacterDefence `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}

type CreateCharacter struct {
	Name     string `binding:"required" json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string ` json:"last_name" query:"last_name" form:"last_name"`
}

type CharacterExternal struct {
	ID             uint            `json:"id"`
	Name           string          `binding:"required" json:"name" query:"name" form:"name"`
	Alias          string          `json:"alias" query:"alias" form:"alias"`
	LastName       string          ` json:"last_name" query:"last_name" form:"last_name"`
	Level          int8            `json:"level" query:"level" form:"level"`
	UserID         uint            ` json:"user_id" query:"user_id" form:"user_id"`
	Attribute      Attribute       `json:"attribute" query:"attribute" form:"attribute"`
	CharacterItem  []CharacterItem `json:"character_item" query:"character_item" form:"character_item"`
	Slot           []Slot          `json:"slot" query:"slot" form:"slot"`
	CharacterBoost CharacterBoost  `json:"character_boost" query:"character_boost" form:"character_boost"`
}

type CharacterUpdate struct {
	Name     string `json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string `json:"last_name" query:"last_name" form:"last_name"`
	Level    int8   `json:"level" query:"level" form:"level"`
}
