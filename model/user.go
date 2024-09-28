package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"unique;not null"`
	Email    *string `gorm:"unique;type:varchar(100)"`
	Password []byte  `gorm:"not null"`
	Admin    bool    `gorm:"default:false"`
}

type UserExternal struct {
	ID       uint   `json:"id"`
	Username string `binding:"required" json:"username" query:"username" form:"username"`
	Admin    bool   `json:"admin" query:"admin" form:"admin"`
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
