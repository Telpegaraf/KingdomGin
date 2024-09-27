package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"unique"`
	Email    *string `gorm:"unique;type:varchar(100)"`
	Password []byte
	Admin    bool   `gorm:"default:false"`
	Token    string `gorm:"type:varchar(180);unique_index" json:"token"`
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
