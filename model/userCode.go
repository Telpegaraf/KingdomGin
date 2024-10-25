package model

import "time"

type UserCode struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Code      string    `gorm:"type:varchar(63);not null"`
	CreatedAt time.Time `gorm:"<-:create"`
}
