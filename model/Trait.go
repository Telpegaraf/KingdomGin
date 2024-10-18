package model

type Trait struct {
	ID          uint    `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string  `gorm:"unique;type:varchar(127);not null"`
	Description string  `gorm:"type:text;not null;"`
	Spells      []Spell `gorm:"many2many:spell_traits;"`
}

type CreateTrait struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateTrait struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type TraitExternal struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
