package database

import "kingdom/model"

func (d *GormDatabase) CreateSlot(slot *model.Slot) error {
	return d.DB.Create(&slot).Error
}
