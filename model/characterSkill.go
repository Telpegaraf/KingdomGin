package model

type CharacterSkill struct {
	ID          uint         `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterID uint         `gorm:"not null;uniqueIndex:idx_character_skill"`
	Name        string       `gorm:"not null;uniqueIndex:idx_character_skill"`
	Mastery     MasteryLevel `gorm:"type:mastery_level;default:None"`
}

type CharacterSkillCreate struct {
	CharacterID uint         `json:"character_id" query:"character_id" binding:"required"`
	Name        string       `json:"name" query:"name" binding:"name"`
	Mastery     MasteryLevel `json:"mastery" query:"mastery" binding:"required" example:"None"`
}

type CharacterSkillUpdate struct {
	Mastery MasteryLevel `json:"mastery" query:"mastery"`
}

type CharacterSkillExternal struct {
	ID          uint         `json:"id" query:"id"`
	CharacterID uint         `json:"character_id" query:"character_id"`
	Name        string       `json:"name" query:"name"`
	Mastery     MasteryLevel `json:"mastery" query:"mastery"`
}
