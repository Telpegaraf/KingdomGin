package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           `gorm:"primaryKey; AUTO_INCREMENT; UNIQUE_INDEX;"`
	Username     string         `gorm:"unique;"`
	Email        string         `gorm:"unique;type:varchar(100)"`
	Password     []byte         `gorm:"not null"`
	Admin        bool           `gorm:"default:false"`
	CreatedAt    time.Time      `gorm:"<-:create"`
	UpdatedAt    time.Time      `gorm:"<-:update"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Verification bool           `gorm:"default:false"`
	Characters   []Character    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserExternal struct {
	ID           uint        `json:"id"`
	Username     string      `binding:"required" json:"username" query:"username" form:"username"`
	Admin        bool        `json:"admin" query:"admin" form:"admin"`
	Email        string      `json:"email"`
	Characters   []Character `json:"characters" query:"characters" form:"characters"`
	Verification bool        `json:"verification" query:"verification" form:"verification"`
}

type CreateUser struct {
	Password string `json:"password,omitempty" query:"password" form:"password" binding:"required"`
	Email    string `binding:"required" json:"email" query:"email" form:"email"`
}

type UserLogin struct {
	Email    string `binding:"required" json:"email" query:"email" form:"email"`
	Password string `binding:"required" json:"password,omitempty" query:"password" form:"password" binding:"required"`
}

type UserUpdateExternal struct {
	Username string  `json:"username" query:"username" form:"username"`
	Email    *string `json:"email" query:"email" form:"email"`
}

type UserPasswordUpdate struct {
	Password string `json:"password,omitempty" query:"password" form:"password" binding:"required"`
}

type UserMessage struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
