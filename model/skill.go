package model

type Skill struct {
	ID          uint         `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string       `gorm:"type:varchar(127);not null;unique"`
	Description string       `gorm:"type:text;not null"`
	Background  []Background `gorm:"many2many:skill_backgrounds;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
