package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CreateFeat Creates New Feat object
func (d *GormDatabase) CreateFeat(feat *model.Feat) error { return d.DB.Create(&feat).Error }

// GetFeatByID Returns Feat object by ID
func (d *GormDatabase) GetFeatByID(id uint) (*model.Feat, error) {
	feat := new(model.Feat)
	err := d.DB.Preload("Traits").Find(&feat, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if feat.ID == id {
		return feat, nil
	}
	return nil, err
}

// GetFeatByName Returns Feat by Name or nil
func (d *GormDatabase) GetFeatByName(name string) (*model.Feat, error) {
	feat := new(model.Feat)
	err := d.DB.Where("name = ?", name).First(&feat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if feat.Name == name {
		return feat, nil
	}
	return nil, err
}

// GetFeats Returns All Feat objects
func (d *GormDatabase) GetFeats() (*[]model.Feat, error) {
	var feats []model.Feat
	err := d.DB.Preload("Traits").Find(&feats).Error
	return &feats, err
}

// DeleteFeat Deletes Feat object by ID
func (d *GormDatabase) DeleteFeat(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.Feat{}).Error
}

// UpdateFeat Updates Feat object
func (d *GormDatabase) UpdateFeat(feat *model.Feat) error {
	return d.DB.Save(&feat).Error
}
