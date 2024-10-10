package model

import (
	"errors"
	"gorm.io/gorm"
)

type Race struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"unique;not null"`
	Description  string `gorm:"not null"`
	HitPoints    int8   `gorm:"not null"`
	Size         Size   `gorm:"type:varchar(31);not null"`
	Speed        int8   `gorm:"not null"`
	AbilityBoost string `gorm:"not null"`
	Language     string
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

type RaceCreate struct {
	Name         string `json:"name" binding:"required" query:"name" form:"name"`
	Description  string `json:"description" binding:"required" query:"description" form:"description"`
	HitPoints    int8   `json:"hit_points" binding:"required" query:"hit_points" form:"hit_points"`
	Size         Size   `json:"size" binding:"required" query:"size" form:"size"`
	Speed        int8   `json:"speed" binding:"required" query:"speed" form:"speed"`
	AbilityBoost string `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	Language     string `json:"language" query:"language" form:"language"`
}