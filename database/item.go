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
