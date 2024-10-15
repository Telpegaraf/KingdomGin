package database

import (
	"kingdom/model"
)

// GetItems Returns all items
func (d *GormDatabase) GetItems() ([]*model.Item, error) {
	var items []*model.Item
	err := d.DB.Where("owner_type", "Armor").Find(&items).Error
	return items, err
}

// GetArmors Returns all armors
func (d *GormDatabase) GetArmors() ([]*model.Armor, error) {
	var armors []*model.Armor
	err := d.DB.Preload("Item").Find(&armors).Error
	return armors, err
}
