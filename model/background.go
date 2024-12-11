package model

type Background struct {
	ID            uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name          string `gorm:"unique;not null"`
	Description   string `gorm:"not null"`
	FeatID        *uint
	FirstSkillID  *uint
	SecondSkillID *uint
	Character     []Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BackgroundCreate struct {
	Name          string `json:"name" binding:"required" query:"name"`
	Description   string `json:"description" binding:"required"`
	FeatID        uint   `json:"feat_id" binding:"required"`
	FirstSkillID  uint   `json:"first_skill_id" binding:"required"`
	SecondSkillID uint   `json:"second_skill_id" binding:"required"`
}

type BackgroundUpdate struct {
	Name          string `json:"name" query:"name"`
	Description   string `json:"description"`
	FeatID        uint   `json:"feat_id"`
	FirstSkillID  uint   `json:"first_skill_id"`
	SecondSkillID uint   `json:"second_skill_id"`
}

type BackgroundExternal struct {
	ID            uint   `json:"id"`
	Name          string `json:"name" query:"name"`
	Description   string `json:"description"`
	FeatID        *uint  `json:"feat_id"`
	FirstSkillID  *uint  `json:"first_skill_id"`
	SecondSkillID *uint  `json:"second_skill_id"`
	//Feat          Feat   `json:"feat"`
}
