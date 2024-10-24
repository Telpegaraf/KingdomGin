package database

import "kingdom/model"

// CreateCharacterInfo create Character Info object
func (d *GormDatabase) CreateCharacterInfo(*model.CharacterInfo) error {
	return d.DB.Create(&model.CharacterInfo{}).Error
}
