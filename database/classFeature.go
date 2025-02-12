package database

import (
	"kingdom/model"
)

// GetClassFeatureByID returns Class Feature by ID
func (d *GormDatabase) GetClassFeatureByID(id uint) (*model.ClassFeature, error) {
	classFeature := new(model.ClassFeature)
	err := d.DB.Preload("SkillFeatures").First(classFeature, id).Error
	if err != nil {
		return nil, err
	}
	return classFeature, nil
}

// GetClassFeatureByClassID returns all Feature for certain class
func (d *GormDatabase) GetClassFeatureByClassID(classID uint) ([]model.ClassFeature, error) {
	var classFeature []model.ClassFeature
	err := d.DB.
		Order("level").
		Preload("SkillFeatures").
		Where("character_class_id = ?", classID).
		Find(&classFeature).Error
	return classFeature, err
}

// GetSkillFeatureByID returns all skills for class for certain level
func (d *GormDatabase) GetSkillFeatureByID(classFeatureID uint) ([]model.SkillFeature, error) {
	var classSkillFeatures []model.SkillFeature
	err := d.DB.Where("class_feature_id = ?", classFeatureID).Find(&classSkillFeatures).Error
	return classSkillFeatures, err
}
