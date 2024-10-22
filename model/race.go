package model

import (
	"errors"
	"gorm.io/gorm"
)

type Race struct {
	ID            uint       `gorm:"primary_key;AUTO_INCREMENT"`
	Name          string     `gorm:"unique;not null"`
	Description   string     `gorm:"not null"`
	HitPoint      uint16     `gorm:"default:6"`
	Size          SquareSize `gorm:"type:square_size;default:Medium"`
	Speed         uint8      `gorm:"not null;default:25"`
	AbilityBoost  uint8      `gorm:"not null;default:2"`
	AttributeFlaw Ability    `gorm:"type:ability"`
	Language      string
	Ancestry      []Ancestry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RaceCreate struct {
	Name          string     `json:"name" binding:"required" query:"name" form:"name"`
	Description   string     `json:"description" binding:"required" query:"description" form:"description"`
	HitPoint      uint16     `json:"hit_points" binding:"required" query:"hit_points" form:"hit_points"`
	Size          SquareSize `json:"size" binding:"required" query:"size" form:"size"`
	Speed         uint8      `json:"speed" binding:"required" query:"speed" form:"speed"`
	AbilityBoost  uint8      `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	AttributeFlaw Ability    `json:"attribute_flaw" query:"attribute_flaw" form:"attribute_flaw"`
	Language      string     `json:"language" query:"language" form:"language"`
}

type RaceUpdate struct {
	Name          string     `json:"name" query:"name" form:"name"`
	Description   string     `json:"description" query:"description" form:"description"`
	HitPoint      uint16     `json:"hit_points" query:"hit_points" form:"hit_points"`
	Size          SquareSize `json:"size" query:"size" form:"size"`
	Speed         uint8      `json:"speed" query:"speed" form:"speed"`
	AbilityBoost  uint8      `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	AttributeFlaw Ability    `json:"attribute_flaw" query:"attribute_flaw" form:"attribute_flaw"`
	Language      string     `json:"language" query:"language" form:"language"`
}

type RaceExternal struct {
	ID            uint       `json:"id" query:"id" form:"id"`
	Name          string     `json:"name" query:"name" form:"name"`
	Description   string     `json:"description" query:"description" form:"description"`
	HitPoint      uint16     `json:"hit_points" query:"hit_points" form:"hit_points"`
	Size          SquareSize `json:"size" query:"size" form:"size"`
	Speed         uint8      `json:"speed" query:"speed" form:"speed"`
	AbilityBoost  uint8      `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	AttributeFlaw Ability    `json:"attribute_flaw" query:"attribute_flaw" form:"attribute_flaw"`
	Language      string     `json:"language" query:"language" form:"language"`
}

func (r *Race) BeforeSave(tx *gorm.DB) (err error) {
	if !isValidSquareSize(r.Size) {
		return errors.New("invalid Square Size vale")
	}
	return
}

func isValidSquareSize(value SquareSize) bool {
	switch value {
	case Tiny, Small, Medium, Large, Huge, Gargantuan:
		return true
	}
	return false
}
