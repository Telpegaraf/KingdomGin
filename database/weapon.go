package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetWeapons Returns all weapons
func (d *GormDatabase) GetWeapons() ([]*model.Weapon, error) {
	var weapons []*model.Weapon
	err := d.DB.Preload("Item").Find(&weapons).Error
	return weapons, err
}

// GetWeaponByID Returns Weapon by ID
func (d *GormDatabase) GetWeaponByID(id uint) (*model.Weapon, error) {
	weapon := new(model.Weapon)
	err := d.DB.Preload("Item").Find(weapon, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if weapon.ID == id {
		return weapon, nil
	}
	return nil, err
}

// CreateWeapon creates Weapon and Item with Owner ID
func (d *GormDatabase) CreateWeapon(weapon *model.Weapon, item *model.Item) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(weapon).Error; err != nil {
			return err
		}
		item.OwnerID = weapon.ID
		if err := tx.Create(item).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWeapon updates Weapon and Item with Owner ID
func (d *GormDatabase) UpdateWeapon(weapon *model.Weapon, item *model.Item) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(weapon).Error; err != nil {
			return err
		}
		if err := tx.Updates(item).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
