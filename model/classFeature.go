package model

type ClassFeature struct {
	ID                uint           `gorm:"primaryKey"`
	CharacterClass    CharacterClass `gorm:"foreignKey:CharacterClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterClassID  uint
	Level             uint8
	IsClassFeat       bool `gorm:"default:false"`
	IsSkillFeat       bool `gorm:"default:false"`
	IsCharacterBoost  bool `gorm:"default:false"`
	IsGeneralFeat     bool `gorm:"default:false"`
	IsSkillIncrease   bool `gorm:"default:false"`
	IsAncestryFeat    bool `gorm:"default:false"`
	WeaponMastery     *MasteryLevel
	ArmorMastery      *MasteryLevel
	PerceptionMastery *MasteryLevel
	FortitudeMastery  *MasteryLevel
	ReflexMastery     *MasteryLevel
	WillMastery       *MasteryLevel
	SkillFeatures     []SkillFeature `gorm:"foreignKey:ClassFeatureID"`
}

type SkillFeature struct {
	ID             uint `gorm:"primaryKey"`
	ClassFeatureID uint
	ClassFeature   ClassFeature `gorm:"foreignKey:ClassFeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name           string       `gorm:"type:varchar(127)"`
	Description    string
}

type ClassFeatureExternal struct {
	ID                uint                    `json:"id"`
	CharacterClassID  uint                    `json:"character_class_id"`
	Level             uint8                   `json:"level"`
	IsClassFeat       bool                    `json:"is_class_feat"`
	IsSkillFeat       bool                    `json:"is_skill_feat"`
	IsCharacterBoost  bool                    `json:"is_character_boost"`
	IsGeneralFeat     bool                    `json:"is_general_feat"`
	IsSkillIncrease   bool                    `json:"is_skill_increase"`
	IsAncestryFeat    bool                    `json:"is_ancestry_feat"`
	WeaponMastery     *MasteryLevel           `json:"weapon_mastery"`
	ArmorMastery      *MasteryLevel           `json:"armor_mastery"`
	PerceptionMastery *MasteryLevel           `json:"perception_mastery"`
	FortitudeMastery  *MasteryLevel           `json:"fortitude_mastery"`
	ReflexMastery     *MasteryLevel           `json:"reflex_mastery"`
	WillMastery       *MasteryLevel           `json:"will_mastery"`
	SkillFeature      []*SkillFeatureExternal `json:"class_skill_feature"`
}

type SkillFeatureExternal struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
