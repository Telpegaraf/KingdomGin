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
	HitPoint          uint16       `gorm:"default:6"`
	TemporaryHitPoint uint16       `gorm:"default:0"`
	Dying             uint8        `gorm:"default:0"`
	Wounded           bool         `gorm:"default:false"`
	Speed             uint8        `gorm:"default:0"`
	CharacterID       uint
}

type CharacterDefenceUpdate struct {
	ArmorClass        uint8        `json:"armorClass"`
	LightArmor        MasteryLevel `json:"light_armor"`
	MediumArmor       MasteryLevel `json:"medium_armor"`
	HeavyArmor        MasteryLevel `json:"heavy_armor"`
	Fortitude         MasteryLevel `json:"fortitude"`
	Reflex            MasteryLevel `json:"reflex"`
	Will              MasteryLevel `json:"will"`
	Perception        MasteryLevel `json:"perception"`
	MaxHitPoint       uint16       `json:"max_hit_point"`
	HitPoint          *uint16      `json:"hit_point"`
	TemporaryHitPoint uint16       `json:"temporary_hit_point"`
	Dying             uint8        `json:"dying"`
	Wounded           bool         `json:"wounded"`
	Speed             uint8        `json:"speed"`
}

type CharacterDefenceExternal struct {
	ID                uint         `json:"id"`
	ArmorClass        uint8        `json:"armorClass"`
	LightArmor        MasteryLevel `json:"light_armor"`
	MediumArmor       MasteryLevel `json:"medium_armor"`
	HeavyArmor        MasteryLevel `json:"heavy_armor"`
	Fortitude         MasteryLevel `json:"fortitude"`
	Reflex            MasteryLevel `json:"reflex"`
	Will              MasteryLevel `json:"will"`
	Perception        MasteryLevel `json:"perception"`
	MaxHitPoint       uint16       `json:"max_hit_point"`
	HitPoint          uint16       `json:"hit_point"`
	TemporaryHitPoint uint16       `json:"temporary_hit_point"`
	Dying             uint8        `json:"dying"`
	Wounded           bool         `json:"wounded"`
	Speed             uint8        `json:"speed"`
	CharacterID       uint         `json:"character_id"`
}
