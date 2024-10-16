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

// CreateArmor creates Armor and Item with Owner ID
func (d *GormDatabase) CreateArmor(armor *model.Armor, item *model.Item) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(armor).Error; err != nil {
			return err
		}
		item.OwnerID = armor.ID
		if err := tx.Create(item).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
