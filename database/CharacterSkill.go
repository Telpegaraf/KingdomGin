package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// CharacterSkillCreate creates Character Skill object
func (d *GormDatabase) CharacterSkillCreate(characterSkill *model.CharacterSkill) error {
	return d.DB.Create(characterSkill).Error
}

// GetCharacterSkillByID returns Character Skill object by ID
func (d *GormDatabase) GetCharacterSkillByID(id uint) (*model.CharacterSkill, error) {
	characterSkill := new(model.CharacterSkill)
	err := d.DB.Find(&characterSkill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return characterSkill, nil
}

// GetCharacterSkillByCharacterID returns Character Skill object by ID
func (d *GormDatabase) GetCharacterSkillByCharacterID(characterId uint, skillID uint) (*model.CharacterSkill, error) {
	characterSkill := new(model.CharacterSkill)
	err := d.DB.Where("character_id = ? AND skill_id = ?", characterId, skillID).First(&characterSkill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return characterSkill, nil
}

// GetCharacterSkills returns all character skill by character ID
func (d *GormDatabase) GetCharacterSkills(id uint) ([]*model.CharacterSkill, error) {
	var characterSkills []*model.CharacterSkill
	err := d.DB.Where("character_id = ?", id).Find(&characterSkills, id).Error
	return characterSkills, err
}

// UpdateCharacterSkill updates character skill object by ID
func (d *GormDatabase) UpdateCharacterSkill(skill *model.CharacterSkill) error {
	return d.DB.Updates(skill).Error
}
