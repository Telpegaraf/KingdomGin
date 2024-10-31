package model

type Skill struct {
	ID               uint         `gorm:"primary_key;AUTO_INCREMENT"`
	Name             string       `gorm:"type:varchar(127);not null;unique"`
	Ability          Ability      `gorm:"type:ability"`
	Description      string       `gorm:"type:text;not null"`
	FirstBackground  []Background `gorm:"foreignKey:FirstSkillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SecondBackground []Background `gorm:"foreignKey:SecondSkillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Feat             []Feat       `gorm:"foreignKey:PrerequisiteSkillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
