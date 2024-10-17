package model

type Character struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null;type:varchar(120)"`
	Alias          string `gorm:"type:varchar(120)"`
	LastName       string `gorm:"type:varchar(120)" json:"last_name"`
	Level          int8   `gorm:"default:1"`
	UserID         uint
	Stat           Attribute       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterItem  []CharacterItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Slot           []Slot          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterBoost CharacterBoost  `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}

type CreateCharacter struct {
	Name     string `binding:"required" json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string ` json:"last_name" query:"last_name" form:"last_name"`
}

type CharacterExternal struct {
	ID       uint   `json:"id"`
	Name     string `binding:"required" json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string ` json:"last_name" query:"last_name" form:"last_name"`
	UserID   uint   ` json:"user_id" query:"user_id" form:"user_id"`
}

type CharacterUpdateExternal struct {
	Name     string `json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string `json:"last_name" query:"last_name" form:"last_name"`
}
