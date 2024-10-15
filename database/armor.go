package database

import "kingdom/model"

// GetArmors Returns all armors
func (d *GormDatabase) GetArmors() ([]*model.Armor, error) {
	var armors []*model.Armor
	err := d.DB.Preload("Item").Find(&armors).Error
	return armors, err
}
