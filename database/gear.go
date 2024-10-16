package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetGears Returns all gears
func (d *GormDatabase) GetGears() ([]*model.Gear, error) {
	var gears []*model.Gear
	err := d.DB.Preload("Item").Find(&gears).Error
	return gears, err
}

// GetGearByID Returns Gear by ID
func (d *GormDatabase) GetGearByID(id uint) (*model.Gear, error) {
	Gear := new(model.Gear)
	err := d.DB.Preload("Item").Find(Gear, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if Gear.ID == id {
		return Gear, nil
	}
	return nil, err
}

// CreateGear creates Gear and Item with Owner ID
func (d *GormDatabase) CreateGear(Gear *model.Gear, item *model.Item) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(Gear).Error; err != nil {
			return err
		}
		item.OwnerID = Gear.ID
		if err := tx.Create(item).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateGear updates Gear and Item with Owner ID
func (d *GormDatabase) UpdateGear(Gear *model.Gear, item *model.Item) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(Gear).Error; err != nil {
			return err
		}
		if err := tx.Model(&item).Select("Level").Updates(item).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
