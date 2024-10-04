package model

import (
	"errors"
	"gorm.io/gorm"
)

type CharacterClass struct {
	ID                uint         `gorm:"primary_key;AUTO_INCREMENT"`
	Name              string       `gorm:"type:varchar(127);unique;not null"`
	Health            int8         `gorm:"not null;default:6"`
	PerceptionMastery MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//FortitudeMastery     MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//ReflexMastery        MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//WillMastery          MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//UnarmedMastery       MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//LightArmorMastery    MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//MediumArmorMastery   MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//HeavyArmorMastery    MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//UnArmedWeaponMastery MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//CommonWeaponMastery  MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	//MartialWeaponMastery MasteryLevel `gorm:"type:mastery"`
}

func (c *CharacterClass) BeforeSave(tx *gorm.DB) error {
	if !isValidMasteryLevel(Health(c.Health)) {
		return errors.New("invalid Health value")
	}
	return nil
}

func isValidMasteryLevel(level Health) bool {
	switch level {
	case Six, Eight, Ten, Twelve:
		return true
	}
	return false
}

type CharacterClassCreate struct {
	Name              string `json:"name" query:"name" form:"name" example:"Fighter"`
	Health            int8   `json:"health" query:"health" form:"health" example:"6" enum:"6,8,10,12"`
	PerceptionMastery string `json:"perception_mastery" query:"perception_mastery" form:"perception_mastery" example:"Train"`
}
