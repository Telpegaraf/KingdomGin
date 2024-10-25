package model

import "time"

type UserCode struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Code      string    `gorm:"type:varchar(63);not null"`
	CreatedAt time.Time `gorm:"<-:create"`
}

type UserCodeVerification struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
