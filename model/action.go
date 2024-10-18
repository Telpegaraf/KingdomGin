package model

type Action struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(127);not null"`

	Spells []Spell `gorm:"many2many:spell_actions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateAction struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAction struct {
	Name string `json:"name"`
}

type ActionExternal struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
