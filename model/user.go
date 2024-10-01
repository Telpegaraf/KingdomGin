package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint    `gorm:"primaryKey"`
	Username   string  `gorm:"unique;not null"`
	Email      *string `gorm:"unique;type:varchar(100)"`
	Password   []byte  `gorm:"not null"`
	Admin      bool    `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Characters []Character    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserExternal struct {
	ID         uint        `json:"id"`
	Username   string      `binding:"required" json:"username" query:"username" form:"username"`
	Admin      bool        `json:"admin" query:"admin" form:"admin"`
	Characters []Character `json:"characters" query:"characters" form:"characters"`
}

type CreateUser struct {
	Username string `binding:"required" json:"username" query:"username" form:"username"`
	Admin    bool   `json:"admin" query:"admin" form:"admin"`
	Password string `json:"password,omitempty" query:"password" form:"password" binding:"required"`
}

type UserLogin struct {
	Username string `binding:"required" json:"username" query:"username" form:"username"`
	Password string `binding:"required" json:"password,omitempty" query:"password" form:"password" binding:"required"`
}
