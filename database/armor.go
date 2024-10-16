package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetArmors Returns all armors
func (d *GormDatabase) GetArmors() ([]*model.Armor, error) {
	var armors []*model.Armor
	err := d.DB.Preload("Item").Find(&armors).Error
	return armors, err
}

// GetArmorByID Returns Armor by ID
func (d *GormDatabase) GetArmorByID(id uint) (*model.Armor, error) {
	armor := new(model.Armor)
	err := d.DB.Preload("Item").Find(armor, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if armor.ID == id {
		return armor, nil
	}
	return nil, err
}
