package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateAncestry creates new Ancestry object
func (d *GormDatabase) CreateAncestry(ancestry *model.Ancestry) error {
	return d.DB.Create(ancestry).Error
}

// GetAncestryByID returns Ancestry by ID
func (d *GormDatabase) GetAncestryByID(id uint) (*model.Ancestry, error) {
	var ancestry model.Ancestry
	err := d.DB.Find(&ancestry, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &ancestry, nil
}

// GetAncestryByName returns Ancestry by Name
func (d *GormDatabase) GetAncestryByName(name string) (*model.Ancestry, error) {
	var ancestry model.Ancestry
	err := d.DB.Where("name = ?", name).First(&ancestry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &ancestry, nil
}

// GetAncestries returns all ancestries
func (d *GormDatabase) GetAncestries() ([]*model.Ancestry, error) {
	var ancestries []*model.Ancestry
	err := d.DB.Find(&ancestries).Error
	return ancestries, err
}

// UpdateAncestry updates ancestry
func (d *GormDatabase) UpdateAncestry(ancestry *model.Ancestry) error {
	return d.DB.Updates(ancestry).Error
}

// DeleteAncestry deletes ancestry
func (d *GormDatabase) DeleteAncestry(id uint) error { return d.DB.Delete(&model.Ancestry{}, id).Error }
