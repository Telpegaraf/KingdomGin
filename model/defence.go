package model

type Defence struct {
	ID                uint   `gorm:"primary_key"`
	ArmorClass        int8   `gorm:"default:10"`
	Unarmed           string `gorm:"type:varchar(20);default:'None';not null"`
	LightArmor        string `gorm:"type:varchar(20);default:'None';not null"`
	MediumArmor       string `gorm:"type:varchar(20);default:'None';not null"`
	HeavyArmor        string `gorm:"type:varchar(20);default:'None';not null"`
	Fortitude         string `gorm:"type:varchar(20);default:'None';not null"`
	Reflex            string `gorm:"type:varchar(20);default:'None';not null"`
	Will              string `gorm:"type:varchar(20);default:'None';not null"`
	Perception        string `gorm:"type:varchar(20);default:'None';not null"`
	MaxHitPoint       int16  `gorm:"default:6"`
	HitPoint          int16  `gorm:"default:6"`
	TemporaryHitPoint int16  `gorm:"default:0"`
	Dying             int8   `gorm:"default:0"`
	Wounded           bool   `gorm:"default:false"`
	Speed             int16  `gorm:"default:0"`
}
