package model

type ClassFeature struct {
	ID                uint           `gorm:"primaryKey"`
	CharacterClass    CharacterClass `gorm:"foreignKey:CharacterClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Level             uint8
	IsClassFeat       bool
	IsSkillFeat       bool
	IsCharacterBoost  bool
	IsGeneralFeat     bool
	IsSkillIncrease   bool
	IsAncestryFeat    bool
	WeaponMastery     MasteryLevel
	ArmorMastery      MasteryLevel
	PerceptionMastery MasteryLevel
	FortitudeMastery  MasteryLevel
	ReflexMastery     MasteryLevel
	WillMastery       MasteryLevel
}

type ClassSkillFeature struct {
	id             uint `gorm:"primaryKey"`
	ClassFeatureID uint
	ClassFeature   ClassFeature `gorm:"foreignKey:ClassFeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name           string       `gorm:"type:varchar(127)"`
	Description    string       `gorm:"type:varchar(127)"`
}
