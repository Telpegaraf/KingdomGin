package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

func (d *GormDatabase) CreateSkill(skill *model.Skill) error {
	return d.DB.Create(skill).Error
}

func (d *GormDatabase) GetSkillByID(id uint) (*model.Skill, error) {
	skill := new(model.Skill)
	err := d.DB.Find(&skill, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if skill.ID == id {
		return skill, err
	}
	return nil, err
}

func (d *GormDatabase) GetSkillByName(name string) (*model.Skill, error) {
	skill := new(model.Skill)
	err := d.DB.Where("name = ?", name).First(skill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if skill.Name == name {
		return skill, err
	}
	return nil, err
}

func (d *GormDatabase) GetSkills() ([]*model.Skill, error) {
	var skills []*model.Skill
	err := d.DB.Find(&skills).Error
	return skills, err
}

func (d *GormDatabase) UpdateSkill(skill *model.Skill) error {
	return d.DB.Updates(skill).Error
}

func (d *GormDatabase) DeleteSkill(id uint) error {
	return d.DB.Where("id = ?", id).
		Delete(&model.Skill{}).Error
}
