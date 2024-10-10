package model

type ImmunityResistanceWeakness struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"unique;not null"`
}
