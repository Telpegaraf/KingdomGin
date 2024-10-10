package model

type Defence struct {
	ID                uint         `gorm:"primary_key"`
	ArmorClass        int8         `gorm:"default:10"`
	Unarmed           MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	LightArmor        MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	MediumArmor       MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	HeavyArmor        MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	Fortitude         MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	Reflex            MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	Will              MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	Perception        MasteryLevel `gorm:"type:varchar(20);default:'None';not null"`
	MaxHitPoint       int16        `gorm:"default:6"`
	HitPoint          int16        `gorm:"default:6"`
	TemporaryHitPoint int16        `gorm:"default:0"`
	Dying             int8         `gorm:"default:0"`
	Wounded           bool         `gorm:"default:false"`
	Speed             int16        `gorm:"default:0"`
}
