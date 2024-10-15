package database

import "kingdom/model"

// GetWeapons Returns all weapons
func (d *GormDatabase) GetWeapons() ([]*model.Weapon, error) {
	var weapons []*model.Weapon
	err := d.DB.Preload("Item").Find(&weapons).Error
	return weapons, err
}
