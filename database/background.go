package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetBackgroundByID Returns Background object by ID
func (d *GormDatabase) GetBackgroundByID(id uint) (*model.Background, error) {
	background := model.Background{}
	err := d.DB.First(&background, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &background, err
}

// GetBackgroundByName Returns Background object by Name
func (d *GormDatabase) GetBackgroundByName(name string) (*model.Background, error) {
	background := model.Background{}
	err := d.DB.Where("name = ?", name).First(&background).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &background, err
}

// GetBackgrounds Returns All Background objects
func (d *GormDatabase) GetBackgrounds() ([]*model.Background, error) {
	var backgrounds []*model.Background
	err := d.DB.Find(&backgrounds).Error
	return backgrounds, err
}

// CreateBackground Creates new Background object
func (d *GormDatabase) CreateBackground(background *model.Background) error {
	return d.DB.Create(background).Error
}

// UpdateBackground Updates Background object
func (d *GormDatabase) UpdateBackground(background *model.Background) error {
	return d.DB.Save(background).Error
}

// DeleteBackground Deletes Background object
func (d *GormDatabase) DeleteBackground(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.Background{}).Error
}
