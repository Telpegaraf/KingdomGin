package database

import "kingdom/model"

// CreateRace creates new Race object
func (d *GormDatabase) CreateRace(race *model.Race) error { return d.DB.Create(&race).Error }
