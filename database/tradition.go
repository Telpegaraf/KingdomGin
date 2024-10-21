package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetTraditionByID returns Tradition by ID
func (d *GormDatabase) GetTraditionByID(id uint) (*model.Tradition, error) {
	tradition := new(model.Tradition)
	err := d.DB.Find(tradition, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if tradition.ID == id {
		return tradition, err
	}
	return nil, err
}

// CreateTradition create new Tradition
func (d *GormDatabase) CreateTradition(tradition *model.Tradition) error {
	return d.DB.Create(tradition).Error
}

// GetTraditions returns all Traditions
func (d *GormDatabase) GetTraditions() ([]*model.Tradition, error) {
	var traditions []*model.Tradition
	err := d.DB.Find(&traditions).Error
	return traditions, err
}

// UpdateTradition updates Tradition
func (d *GormDatabase) UpdateTradition(tradition *model.Tradition) error {
	return d.DB.Save(tradition).Error
}

// DeleteTradition deletes Tradition
func (d *GormDatabase) DeleteTradition(id uint) error {
	return d.DB.Where("id = ?", id).
		Delete(&model.Tradition{}).Error
}
