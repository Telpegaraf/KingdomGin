package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetItems Returns all items
func (d *GormDatabase) GetItems() ([]*model.Item, error) {
	var items []*model.Item
	err := d.DB.Find(&items).Error
	return items, err
}

// GetItemByID Returns Item by ID
func (d *GormDatabase) GetItemByID(id uint) (*model.Item, error) {
	item := new(model.Item)
	err := d.DB.Find(item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if item.ID == id {
		return item, err
	}
	return nil, err
}

// DeleteItem Deletes item by ID
func (d *GormDatabase) DeleteItem(id uint, ownerType string, ownerID uint) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		switch ownerType {
		case "armors":
			err := d.DB.Delete(&model.Armor{}, "id = ?", ownerID).Error
			if err != nil {
				return err
			}
		case "weapons":
			err := d.DB.Delete(&model.Weapon{}, "id = ?", ownerID).Error
			if err != nil {
				return err
			}
		case "gears":
			err := d.DB.Delete(&model.Gear{}, "id = ?", ownerID).Error
			if err != nil {
				return err
			}
		}
		err := d.DB.Delete(&model.Item{}, "id = ?", id).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
