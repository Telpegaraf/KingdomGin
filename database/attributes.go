package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateAttribute creates new Attributes object, linked with Character
func (d *GormDatabase) CreateAttribute(stat *model.Attributes) error { return d.DB.Create(stat).Error }

// GetAttributeByID returns Attribute object by ID
func (d *GormDatabase) GetAttributeByID(id uint) (*model.Attributes, error) {
	attribute := &model.Attributes{}
	err := d.DB.Find(attribute, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if attribute.ID == id {
		return attribute, nil
	}
	return nil, err
}

// GetAllAttributes returns all Attributes objects
func (d *GormDatabase) GetAllAttributes() (*[]model.Attributes, error) {
	var attributes []model.Attributes
	err := d.DB.Find(&attributes).Error
	return &attributes, err
}
