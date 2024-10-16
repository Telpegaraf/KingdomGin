package database

import (
	"gorm.io/gorm"
	"kingdom/model"
)

// GetItems Returns all items
func (d *GormDatabase) GetItems() ([]*model.Item, error) {
	var items []*model.Item
	err := d.DB.Find(&items).Error
	return items, err
}

// DeleteItem Deletes item by ID
func (d *GormDatabase) DeleteItem(id uint) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		err := d.DB.Delete(&model.Item{}, "id = ?", id).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
