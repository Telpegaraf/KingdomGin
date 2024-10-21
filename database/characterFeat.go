package database

import "kingdom/model"

// CreateCharacterFeat creates Character Feat
func (d *GormDatabase) CreateCharacterFeat(characterFeat *model.CharacterFeat) error {
	return d.DB.Create(characterFeat).Error
}
