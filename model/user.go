package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `gorm:"primaryKey; AUTO_INCREMENT; UNIQUE_INDEX;"`
	TgID       uint           `gorm:"unique; UNIQUE_INDEX;"`
	Username   string         `gorm:"unique;"`
	Admin      bool           `gorm:"default:false"`
	CreatedAt  time.Time      `gorm:"<-:create"`
	UpdatedAt  time.Time      `gorm:"<-:update"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Characters []Character    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserExternal struct {
	ID         uint        `json:"id"`
	TgID       uint        `json:"tg_id"`
	Username   string      `binding:"required" json:"username" query:"username" form:"username"`
	Admin      bool        `json:"admin" query:"admin" form:"admin"`
	Characters []Character `json:"characters" query:"characters" form:"characters"`
}

type CreateUser struct {
	TgID uint `json:"tg_id"`
}
