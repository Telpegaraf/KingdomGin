package database

import (
	"kingdom/model"
)

// GetItems Returns all items
func (d *GormDatabase) GetItems() ([]*model.Item, error) {
	var items []*model.Item
	err := d.DB.Find(&items).Error
	return items, err
}

// GetArmors Returns all armors
func (d *GormDatabase) GetArmors() ([]*model.Armor, error) {
	var armors []*model.Armor
	err := d.DB.Preload("Item").Find(&armors).Error
	return armors, err
}

// GetWeapons Returns all weapons
func (d *GormDatabase) GetWeapons() ([]*model.Weapon, error) {
	var weapons []*model.Weapon
	err := d.DB.Preload("Item").Find(&weapons).Error
	return weapons, err
}
