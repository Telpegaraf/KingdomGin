package model

type CharacterDefence struct {
	ID                uint         `gorm:"primary_key"`
	ArmorClass        uint8        `gorm:"default:10"`
	Unarmed           MasteryLevel `gorm:"type:mastery_level;default:None"`
	LightArmor        MasteryLevel `gorm:"type:mastery_level;default:None"`
	MediumArmor       MasteryLevel `gorm:"type:mastery_level;default:None"`
	HeavyArmor        MasteryLevel `gorm:"type:mastery_level;default:None"`
	Fortitude         MasteryLevel `gorm:"type:mastery_level;default:None"`
	Reflex            MasteryLevel `gorm:"type:mastery_level;default:None"`
	Will              MasteryLevel `gorm:"type:mastery_level;default:None"`
	Perception        MasteryLevel `gorm:"type:mastery_level;default:None"`
	MaxHitPoint       uint16       `gorm:"default:6"`
	HitPoint          int16        `gorm:"default:6"`
	TemporaryHitPoint uint16       `gorm:"default:0"`
	Dying             uint8        `gorm:"default:0"`
	Wounded           bool         `gorm:"default:false"`
	Speed             uint8        `gorm:"default:0"`
	CharacterID       uint
}
