package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetCharacterDefenceByID returns Character Defence object by ID
func (d *GormDatabase) GetCharacterDefenceByID(id uint) (*model.CharacterDefence, error) {
	characterDefence := new(model.CharacterDefence)
	err := d.DB.Find(characterDefence, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return characterDefence, nil
}

// UpdateCharacterDefence updates Character Defence object
func (d *GormDatabase) UpdateCharacterDefence(defence *model.CharacterDefence) error {
	return d.DB.Model(&defence).
		Select("armor_class", "hit_point", "temporary_hit_point", "wounded", "speed").
		Updates(defence).Error
}
