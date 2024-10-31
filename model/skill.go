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

type SkillCreate struct {
	Name        string  `json:"name" query:"name"`
	Ability     Ability `json:"ability" query:"ability"`
	Description string  `json:"description" query:"description"`
}

type SkillUpdate struct {
	Name        string  `json:"name" query:"name"`
	Ability     Ability `json:"ability" query:"ability"`
	Description string  `json:"description" query:"description"`
}

type SkillExternal struct {
	ID          uint    `json:"id" query:"id"`
	Name        string  `json:"name" query:"name"`
	Ability     Ability `json:"ability" query:"ability"`
	Description string  `json:"description" query:"description"`
}
