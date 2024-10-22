package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateCharacterBoost creates new Boost object, linked with Character
func (d *GormDatabase) CreateCharacterBoost(stat *model.CharacterBoost) error {
	return d.DB.Create(stat).Error
}

// GetCharacterBoostByID returns Boost object by ID
func (d *GormDatabase) GetCharacterBoostByID(characterID uint) (*model.CharacterBoost, error) {
	characterBoost := &model.CharacterBoost{}
	err := d.DB.Where("character_id = ?", characterID).First(characterBoost).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return characterBoost, nil
}

// UpdateCharacterBoost updates Boost object
func (d *GormDatabase) UpdateCharacterBoost(characterBoost *model.CharacterBoost) error {
	return d.DB.Model(characterBoost).
		Select("strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma", "free").
		Updates(characterBoost).Error
}
