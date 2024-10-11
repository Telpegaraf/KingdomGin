package model

type CharacterClass struct {
	ID                   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name                 string `gorm:"type:varchar(127);unique;not null"`
	Health               int8   `gorm:"not null;default:6"`
	PerceptionMastery    string `gorm:"type:varchar(20);default:'None';not null"`
	FortitudeMastery     string `gorm:"type:varchar(20);default:'None';not null"`
	ReflexMastery        string `gorm:"type:varchar(20);default:'None';not null"`
	WillMastery          string `gorm:"type:varchar(20);default:'None';not null"`
	UnarmedMastery       string `gorm:"type:varchar(20);default:'None';not null"`
	LightArmorMastery    string `gorm:"type:varchar(20);default:'None';not null"`
	MediumArmorMastery   string `gorm:"type:varchar(20);default:'None';not null"`
	HeavyArmorMastery    string `gorm:"type:varchar(20);default:'None';not null"`
	UnArmedWeaponMastery string `gorm:"type:varchar(20);default:'None';not null"`
	CommonWeaponMastery  string `gorm:"type:varchar(20);default:'None';not null"`
	MartialWeaponMastery string `gorm:"type:varchar(20);default:'None';not null"`
}

type CharacterClassCreate struct {
	Name              string `json:"name" query:"name" form:"name" example:"Fighter"`
	Health            int8   `json:"health" query:"health" form:"health" example:"6" enum:"6,8,10,12"`
	PerceptionMastery string `json:"perception_mastery" query:"perception_mastery" form:"perception_mastery" example:"Train"`
}
