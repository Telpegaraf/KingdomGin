package model

type CharacterCondition struct {
	ID          uint `gorm:"primary_key"`
	ConditionID uint
	Count       int8
}
