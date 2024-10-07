package database

import "kingdom/model"

// CreateStat creates new Stat object, linked with Character
func (d *GormDatabase) CreateStat(stat *model.Stat) error { return d.DB.Create(stat).Error }
