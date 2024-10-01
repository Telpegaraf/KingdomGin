package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetGodByID returns God by ID
func (d *GormDatabase) GetGodByID(id uint) (*model.God, error) {
	god := new(model.God)
	err := d.DB.Preload("Domains").Find(god, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if god.ID == id {
		return god, err
	}
	return nil, err
}

// CreateGod create new God
func (d *GormDatabase) CreateGod(god *model.God) error { return d.DB.Create(god).Error }
