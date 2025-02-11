package model

type CharacterClass struct {
	ID            uint         `gorm:"primary_key;AUTO_INCREMENT"`
	Name          string       `gorm:"type:varchar(127);unique;not null"`
	HitPoint      uint16       `gorm:"not null;default:6"`
	Perception    MasteryLevel `gorm:"type:mastery_level;default:None"`
	Fortitude     MasteryLevel `gorm:"type:mastery_level;default:None"`
	Reflex        MasteryLevel `gorm:"type:mastery_level;default:None"`
	Will          MasteryLevel `gorm:"type:mastery_level;default:None"`
	UnarmedArmor  MasteryLevel `gorm:"type:mastery_level;default:None"`
	LightArmor    MasteryLevel `gorm:"type:mastery_level;default:None"`
	MediumArmor   MasteryLevel `gorm:"type:mastery_level;default:None"`
	HeavyArmor    MasteryLevel `gorm:"type:mastery_level;default:None"`
	UnArmedWeapon MasteryLevel `gorm:"type:mastery_level;default:None"`
	CommonWeapon  MasteryLevel `gorm:"type:mastery_level;default:None"`
	MartialWeapon MasteryLevel `gorm:"type:mastery_level;default:None"`
	TraditionID   *uint

	//Character Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CharacterClassCreate struct {
	Name          string       `json:"name" query:"name" form:"name" example:"Fighter"`
	HitPoint      uint16       `json:"health" query:"health" form:"health" example:"6" enum:"6,8,10,12"`
	Perception    MasteryLevel `json:"perception" query:"perception" form:"perception" example:"Train"`
	Fortitude     MasteryLevel `json:"fortitude" query:"fortitude" form:"fortitude" example:"Train"`
	Reflex        MasteryLevel `json:"reflex" query:"reflex" form:"reflex" example:"Train"`
	Will          MasteryLevel `json:"will" query:"will" form:"will" example:"Train"`
	UnarmedArmor  MasteryLevel `json:"unarmed_armor" query:"unarmed_armor" form:"unarmed_armor" example:"Train"`
	LightArmor    MasteryLevel `json:"light_armor" query:"light_armor" form:"light_armor" example:"Train"`
	MediumArmor   MasteryLevel `json:"medium_armor" query:"medium_armor" example:"Train"`
	HeavyArmor    MasteryLevel `json:"heavy_armor" query:"heavy_armor" example:"Train"`
	UnArmedWeapon MasteryLevel `json:"un_armed_weapon" query:"un_armed_weapon" example:"Train"`
	CommonWeapon  MasteryLevel `json:"common_weapon" query:"common_weapon" example:"Train"`
	MartialWeapon MasteryLevel `json:"martial_weapon" query:"martial_weapon" example:"Train"`
	TraditionID   uint         `json:"tradition_id" query:"tradition_id"`
}

type CharacterClassUpdate struct {
	Name          string       `json:"name" query:"name" form:"name" example:"Fighter"`
	HitPoint      uint16       `json:"health" query:"health" form:"health" example:"6" enum:"6,8,10,12"`
	Perception    MasteryLevel `json:"perception" query:"perception" form:"perception" example:"Train"`
	Fortitude     MasteryLevel `json:"fortitude" query:"fortitude" form:"fortitude" example:"Train"`
	Reflex        MasteryLevel `json:"reflex" query:"reflex" form:"reflex" example:"Train"`
	Will          MasteryLevel `json:"will" query:"will" form:"will" example:"Train"`
	UnarmedArmor  MasteryLevel `json:"unarmed_armor" query:"unarmed_armor" form:"unarmed_armor" example:"Train"`
	LightArmor    MasteryLevel `json:"light_armor" query:"light_armor" form:"light_armor" example:"Train"`
	MediumArmor   MasteryLevel `json:"medium_armor" query:"medium_armor" example:"Train"`
	HeavyArmor    MasteryLevel `json:"heavy_armor" query:"heavy_armor" example:"Train"`
	UnArmedWeapon MasteryLevel `json:"un_armed_weapon" query:"un_armed_weapon" example:"Train"`
	CommonWeapon  MasteryLevel `json:"common_weapon" query:"common_weapon" example:"Train"`
	MartialWeapon MasteryLevel `json:"martial_weapon" query:"martial_weapon" example:"Train"`
	TraditionID   uint         `json:"tradition_id" query:"tradition_id"`
}

type CharacterClassExternal struct {
	ID            uint         `json:"id" query:"id" form:"id"`
	Name          string       `json:"name" query:"name" form:"name" example:"Fighter"`
	HitPoint      uint16       `json:"health" query:"health" form:"health" example:"6" enum:"6,8,10,12"`
	Perception    MasteryLevel `json:"perception" query:"perception" form:"perception" example:"Train"`
	Fortitude     MasteryLevel `json:"fortitude" query:"fortitude" form:"fortitude" example:"Train"`
	Reflex        MasteryLevel `json:"reflex" query:"reflex" form:"reflex" example:"Train"`
	Will          MasteryLevel `json:"will" query:"will" form:"will" example:"Train"`
	UnarmedArmor  MasteryLevel `json:"unarmed_armor" query:"unarmed_armor" form:"unarmed_armor" example:"Train"`
	LightArmor    MasteryLevel `json:"light_armor" query:"light_armor" form:"light_armor" example:"Train"`
	MediumArmor   MasteryLevel `json:"medium_armor" query:"medium_armor" example:"Train"`
	HeavyArmor    MasteryLevel `json:"heavy_armor" query:"heavy_armor" example:"Train"`
	UnArmedWeapon MasteryLevel `json:"un_armed_weapon" query:"un_armed_weapon" example:"Train"`
	CommonWeapon  MasteryLevel `json:"common_weapon" query:"common_weapon" example:"Train"`
	MartialWeapon MasteryLevel `json:"martial_weapon" query:"martial_weapon" example:"Train"`
	TraditionID   *uint        `json:"tradition_id" query:"tradition_id"`
}
