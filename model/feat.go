package model

type Feat struct {
	ID                  uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name                string `gorm:"unique;not null"`
	Description         string
	Level               uint8 `gorm:"default:1;not null"`
	PrerequisiteSkillID *uint
	PrerequisiteMastery MasteryLevel    `gorm:"type:mastery_level"`
	Rarity              Rarity          `gorm:"type:rarity;default:Common"`
	Background          []Background    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterFeat       []CharacterFeat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Traits              []Trait         `gorm:"many2many:feat_traits;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}

type CreateFeat struct {
	Name                string       `json:"name" binding:"required" query:"name"`
	Description         string       `json:"description" query:"description"`
	Level               uint8        `json:"level" query:"level"`
	PrerequisiteSkillID *uint        `json:"prerequisite_skill_id" query:"prerequisite_skill_id"`
	PrerequisiteMastery MasteryLevel `gorm:"type:mastery_level"`
}

type UpdateFeat struct {
	Name                string       `json:"name" query:"name"`
	Description         string       `json:"description" query:"description"`
	Level               uint8        `json:"level" query:"level"`
	PrerequisiteSkillID *uint        `json:"prerequisite_skill_id" query:"prerequisite_skill_id"`
	PrerequisiteMastery MasteryLevel `gorm:"type:mastery_level"`
}

type FeatExternal struct {
	ID                  uint         `json:"id" query:"id" form:"id"`
	Name                string       `json:"name" query:"name"`
	Description         string       `json:"description" query:"description"`
	Level               uint8        `json:"level" query:"level"`
	PrerequisiteSkillID *uint        `json:"prerequisite_skill_id" query:"prerequisite_skill_id"`
	PrerequisiteMastery MasteryLevel `gorm:"type:mastery_level"`
}
