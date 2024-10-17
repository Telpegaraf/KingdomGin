package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateSlot creates new slot linked with character
func (d *GormDatabase) CreateSlot(slot *model.Slot) error {
	return d.DB.Create(&slot).Error
}

// GetSlotByID returns slot by ID
func (d *GormDatabase) GetSlotByID(id uint) (*model.Slot, error) {
	slot := new(model.Slot)
	err := d.DB.Find(&slot, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if slot.ID == id {
		return slot, nil
	}
	return nil, err
}

// UpdateSlot updates slot
func (d *GormDatabase) UpdateSlot(slot *model.Slot) error {
	return d.DB.Updates(slot).Error
}
