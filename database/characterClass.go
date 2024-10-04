package database

import "kingdom/model"

// CreateCharacterClass creates new character class
func (d *GormDatabase) CreateCharacterClass(characterClass *model.CharacterClass) error {
	return d.DB.Create(characterClass).Error
}
