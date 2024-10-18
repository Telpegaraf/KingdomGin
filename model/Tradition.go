package model

type Tradition struct {
	ID             uint           `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string         `gorm:"unique;type:varchar(127)"`
	Description    string         `gorm:"type:text"`
	Spells         []Spell        `gorm:"many2many:spell_traditions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CharacterClass CharacterClass `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateTradition struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTradition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TraditionExternal struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
