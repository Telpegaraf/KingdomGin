package database

import "kingdom/model"

// CreateCharacterInfo create Character Info object
func (d *GormDatabase) CreateCharacterInfo(characterInfo *model.CharacterInfo) error {
	return d.DB.Create(characterInfo).Error
}
