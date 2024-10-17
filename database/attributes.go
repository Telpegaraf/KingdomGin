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

// UpdateAttribute updates Attribute object
func (d *GormDatabase) UpdateAttribute(attributes *model.Attribute) error {
	return d.DB.Model(attributes).
		Select("strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma").
		Updates(attributes).Error
}
