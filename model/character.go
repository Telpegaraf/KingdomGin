package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name     string `gorm:"not null;type:varchar(120)"`
	Alias    string `gorm:"type:varchar(120)"`
	LastName string `gorm:"type:varchar(120)"`
}

type CreateCharacter struct {
	Name     string `binding:"required" json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string ` json:"last_name" query:"last_name" form:"last_name"`
}

type CharacterExternal struct {
	Name     string `binding:"required" json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string ` json:"last_name" query:"last_name" form:"last_name"`
}
