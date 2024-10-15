package database

import "kingdom/model"

// GetGears Returns all gears
func (d *GormDatabase) GetGears() ([]*model.Gear, error) {
	var gears []*model.Gear
	err := d.DB.Preload("Item").Find(&gears).Error
	return gears, err
}
