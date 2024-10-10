package model

import (
	"errors"
	"gorm.io/gorm"
)

type CharacterClass struct {
	ID                   uint         `gorm:"primary_key;AUTO_INCREMENT"`
	Name                 string       `gorm:"type:varchar(127);unique;not null"`
	Health               int8         `gorm:"not null;default:6"`
	PerceptionMastery    MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	FortitudeMastery     MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	ReflexMastery        MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	WillMastery          MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	UnarmedMastery       MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	LightArmorMastery    MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	MediumArmorMastery   MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	HeavyArmorMastery    MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	UnArmedWeaponMastery MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	CommonWeaponMastery  MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	MartialWeaponMastery MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
}

func (c *CharacterClass) BeforeSave(tx *gorm.DB) error {
	if !isValidHealth(Health(c.Health)) {
		return errors.New("invalid Health value")
	}
	mysteries := []MasteryLevel{
		c.PerceptionMastery,
		c.FortitudeMastery,
		c.ReflexMastery,
		c.WillMastery,
		c.UnarmedMastery,
		c.LightArmorMastery,
		c.MediumArmorMastery,
		c.HeavyArmorMastery,
		c.UnArmedWeaponMastery,
		c.CommonWeaponMastery,
		c.MartialWeaponMastery,
	}
	masteryNames := []string{
		"PerceptionMastery",
		"FortitudeMastery",
		"ReflexMastery",
		"WillMastery",
		"UnarmedMastery",
		"LightArmorMastery",
		"MediumArmorMastery",
		"HeavyArmorMastery",
		"UnArmedWeaponMastery",
		"CommonWeaponMastery",
		"MartialWeaponMastery",
	}
	for i, mastery := range mysteries {
		if !isValidMasteryLevel(mastery) {
			return errors.New("invalid " + masteryNames[i] + " value")
		}
	}
	return nil
}

func isValidHealth(level Health) bool {
	switch level {
	case Six, Eight, Ten, Twelve:
		return true
	}
	return false
}

func isValidMasteryLevel(level MasteryLevel) bool {
	switch level {
	case None, Train, Expert, Master, Legend:
		return true
	}
	return false
}

type CharacterClassCreate struct {
	Name              string `json:"name" query:"name" form:"name" example:"Fighter"`
	Health            int8   `json:"health" query:"health" form:"health" example:"6" enum:"6,8,10,12"`
	PerceptionMastery string `json:"perception_mastery" query:"perception_mastery" form:"perception_mastery" example:"Train"`
}
