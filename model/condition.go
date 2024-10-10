package model

type Condition struct {
	ID                 uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name               string `gorm:"unique; not null"`
	Description        string `gorm:"type:text"`
	CharacterCondition []CharacterCondition
}
