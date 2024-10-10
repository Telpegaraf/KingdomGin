package database

import "kingdom/model"

// CreateAttribute creates new Attributes object, linked with Character
func (d *GormDatabase) CreateAttribute(stat *model.Attributes) error { return d.DB.Create(stat).Error }
