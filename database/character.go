package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateCharacter create and returns new Character
func (d *GormDatabase) CreateCharacter(character *model.Character) error {
	return d.DB.Create(character).Error
}

// GetCharacterByID returns Character by ID
func (d *GormDatabase) GetCharacterByID(id uint) (*model.Character, error) {
	character := new(model.Character)
	err := d.DB.
		Preload("CharacterClass").
		Preload("Race").
		Preload("Ancestry").
		Preload("Background").
		Preload("Attribute").
		Preload("CharacterSkill").
		Preload("CharacterInfo").
		Find(character, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if character.ID == id {
		return character, err
	}
	return nil, err
}

// GetCharacters returns all characters
func (d *GormDatabase) GetCharacters(id uint) ([]*model.Character, error) {
	var characters []*model.Character
	err := d.DB.
		Preload("CharacterClass").
		Preload("Race").
		Where(&model.Character{UserID: id}).
		Find(&characters).Error
	return characters, err
}

// UpdateCharacter updates character by its id
func (d *GormDatabase) UpdateCharacter(character *model.Character) error {
	return d.DB.Save(character).Error
}

// DeleteCharacterByID deletes character by its id
func (d *GormDatabase) DeleteCharacterByID(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.Character{}, id).Error
}

func (d *GormDatabase) UpdateHitPoint(defence *model.CharacterDefence) error {
	return d.DB.Model(&defence).Select("max_hit_points").Updates(defence).Error
}
