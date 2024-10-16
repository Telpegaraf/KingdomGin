package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateCharacterItem create new character item
func (d *GormDatabase) CreateCharacterItem(characterItem *model.CharacterItem) error {
	return d.DB.Create(characterItem).Error
}

// GetCharacterItem get character item by ID
func (d *GormDatabase) GetCharacterItem(id uint) (*model.CharacterItem, error) {
	characterItem := new(model.CharacterItem)
	err := d.DB.Preload("Character").Preload("Item").Find(characterItem, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if characterItem.ID == id {
		return characterItem, nil
	}
	return nil, err
}

// GetCharacterItems get character items
func (d *GormDatabase) GetCharacterItems() ([]*model.CharacterItem, error) {
	var characterItems []*model.CharacterItem
	err := d.DB.Preload("Character").Preload("Item").Find(&characterItems).Error
	return characterItems, err
}

// UpdateCharacterItem updates character item by ID
func (d *GormDatabase) UpdateCharacterItem(item *model.CharacterItem) error {
	return d.DB.Updates(item).Error
}

// DeleteCharacterItem deletes character item by ID
func (d *GormDatabase) DeleteCharacterItem(id uint) error {
	return d.DB.Delete(&model.CharacterItem{}, id).Error
}
