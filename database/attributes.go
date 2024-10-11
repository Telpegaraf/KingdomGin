package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateAttribute creates new Attribute object, linked with Character
func (d *GormDatabase) CreateAttribute(stat *model.Attribute) error { return d.DB.Create(stat).Error }

// GetAttributeByID returns Attribute object by ID
func (d *GormDatabase) GetAttributeByID(id uint) (*model.Attribute, error) {
	attribute := &model.Attribute{}
	err := d.DB.Find(attribute, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if attribute.ID == id {
		return attribute, nil
	}
	return nil, err
}

// GetAttributes returns all Attribute objects
func (d *GormDatabase) GetAttributes() (*[]model.Attribute, error) {
	var attributes []model.Attribute
	err := d.DB.Find(&attributes).Error
	return &attributes, err
}

// DeleteAttribute deletes Attribute object
func (d *GormDatabase) DeleteAttribute(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.Attribute{}).Error
}

// UpdateAttribute updates Attribute object
func (d *GormDatabase) UpdateAttribute(attributes *model.Attribute) error {
	return d.DB.Save(attributes).Error
}
