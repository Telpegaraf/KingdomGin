package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetTraitByID returns Trait by ID
func (d *GormDatabase) GetTraitByID(id uint) (*model.Trait, error) {
	trait := new(model.Trait)
	err := d.DB.Find(trait, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if trait.ID == id {
		return trait, nil
	}
	return nil, err
}

// GetTraitByName Returns Trait by Name or nil
func (d *GormDatabase) GetTraitByName(name string) (*model.Trait, error) {
	trait := new(model.Trait)
	err := d.DB.Where("name = ?", name).Find(trait).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if trait.Name == name {
		return trait, nil
	}
	return nil, err
}

// CreateTrait create new Trait
func (d *GormDatabase) CreateTrait(trait *model.Trait) error { return d.DB.Create(trait).Error }

// GetTraits returns all Traits
func (d *GormDatabase) GetTraits() ([]*model.Trait, error) {
	var traits []*model.Trait
	err := d.DB.Find(&traits).Error
	return traits, err
}

// UpdateTrait updates Trait
func (d *GormDatabase) UpdateTrait(trait *model.Trait) error { return d.DB.Save(trait).Error }

// DeleteTrait deletes Trait
func (d *GormDatabase) DeleteTrait(id uint) error {
	return d.DB.Where("id = ?", id).
		Delete(&model.Trait{}).Error
}
