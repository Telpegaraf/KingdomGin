package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateCharacterClass creates new character class
func (d *GormDatabase) CreateCharacterClass(characterClass *model.CharacterClass) error {
	return d.DB.Create(characterClass).Error
}

// GetCharacterClassByID returns Character Class by ID
func (d *GormDatabase) GetCharacterClassByID(id uint) (*model.CharacterClass, error) {
	characterClass := &model.CharacterClass{}
	err := d.DB.Find(characterClass, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if characterClass.ID == id {
		return characterClass, nil
	}
	return nil, err
}

// GetCharacterClassByName returns Character Class by Name
func (d *GormDatabase) GetCharacterClassByName(name string) (*model.CharacterClass, error) {
	characterClass := &model.CharacterClass{}
	err := d.DB.
		Order("Name ASC").
		First(characterClass).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if characterClass.Name == name {
		return characterClass, err
	}
	return nil, err
}

// GetCharacterClasses returns all Character classes
func (d *GormDatabase) GetCharacterClasses() ([]*model.CharacterClass, error) {
	var characterClass []*model.CharacterClass
	err := d.DB.Order("Name ASC").Find(&characterClass).Error
	return characterClass, err
}

// DeleteCharacterClass deletes Character Class
func (d *GormDatabase) DeleteCharacterClass(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.CharacterClass{}).Error
}

// UpdateCharacterClass updates Character Class
func (d *GormDatabase) UpdateCharacterClass(class *model.CharacterClass) error {
	return d.DB.Save(class).Error
}
