package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetCharacterByID returns Character for the ID
func (d *GormDatabase) GetCharacterByID(id uint) (*model.Character, error) {
	character := new(model.Character)
	err := d.DB.Find(character, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if character.ID == id {
		return character, err
	}
	return nil, err
}

// CreateCharacter create and returns new Character
func (d *GormDatabase) CreateCharacter(character *model.Character) error {
	return d.DB.Create(character).Error
}
