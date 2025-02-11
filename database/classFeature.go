package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetClassFeatureByID returns Class Feature by ID
func (d *GormDatabase) GetClassFeatureByID(id uint) (*model.ClassFeature, error) {
	classFeature := model.ClassFeature{}
	err := d.DB.First(&classFeature, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &classFeature, err
}

// GetAllClassSkillFeatureByID returns all skills for class for certain level
func (d *GormDatabase) GetAllClassSkillFeatureByID(classFeatureID uint) ([]model.ClassSkillFeature, error) {
	var classSkillFeatures []model.ClassSkillFeature
	err := d.DB.Where("class_feature_id = ?", classFeatureID).Find(&classSkillFeatures).Error
	return classSkillFeatures, err
}
