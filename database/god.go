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

// GetAllGod returns all Gods
func (d *GormDatabase) GetAllGod() ([]*model.God, error) {
	var gods []*model.God
	err := d.DB.Preload("Domains").Find(&gods).Error
	return gods, err
}

// DeleteGodByID deletes God by ID
func (d *GormDatabase) DeleteGodByID(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.God{}).Error
}

// UpdateGod updates God
func (d *GormDatabase) UpdateGod(god *model.God) error {
	return d.DB.Save(god).Error
}
