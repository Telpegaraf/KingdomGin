package database

import "kingdom/model"

// CreateCharacterInfo create Character Info object
func (d *GormDatabase) CreateCharacterInfo(characterInfo *model.CharacterInfo) error {
	return d.DB.Create(characterInfo).Error
}

// GetCharacterInfoByID returns Character Info by ID
func (d *GormDatabase) GetCharacterInfoByID(characterID uint) (*model.CharacterInfo, error) {
	var characterInfo model.CharacterInfo
	err := d.DB.Where("character_id = ?", characterID).First(&characterInfo).Error
	return &characterInfo, err
}

// UpdateCharacterInfo updates character info object
func (d *GormDatabase) UpdateCharacterInfo(characterInfo *model.CharacterInfo) error {
	return d.DB.Updates(characterInfo).Error
}
