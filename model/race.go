package model

import (
	"errors"
	"gorm.io/gorm"
)

type Race struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"unique;not null"`
	HitPoints    int8   `gorm:"not null"`
	Size         Size   `gorm:"type:varchar(31);not null"`
	Speed        int8   `gorm:"not null"`
	AbilityBoost string `gorm:"not null"`
	Language     string `gorm:"not null"`
	Ancestry     []Ancestry
}

func (r Race) BeforeSave(tx *gorm.DB) error {
	if !isValidSize(r.Size) {
		return errors.New("invalid size")
	}
	return nil
}

func isValidSize(size Size) bool {
	switch size {
	case Tiny, Small, Medium, Large, Huge:
		return true
	}
	return false
}
