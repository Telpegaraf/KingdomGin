package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetSpellByID returns Spell by ID
func (d *GormDatabase) GetSpellByID(id uint) (*model.Spell, error) {
	spell := new(model.Spell)
	err := d.DB.Preload("Tradition").Find(spell, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if spell.ID == id {
		return spell, err
	}
	return nil, err
}

// GetSpellByName returns Spell by name
func (d *GormDatabase) GetSpellByName(name string) (*model.Spell, error) {
	spell := new(model.Spell)
	err := d.DB.Where("name = ?", name).First(spell).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if spell.Name == name {
		return spell, err
	}
	return nil, err
}

// CreateSpell create new Spell
func (d *GormDatabase) CreateSpell(spell *model.Spell) error {
	return d.DB.Create(spell).Error
}

// GetSpells returns all Spells
func (d *GormDatabase) GetSpells(limit int, offset int) ([]*model.Spell, error) {
	var spells []*model.Spell
	if limit == 0 {
		err := d.DB.Find(&spells).Error
		return spells, err
	}
	err := d.DB.Limit(limit).Offset(offset).First(&spells).Error
	return spells, err
}

// UpdateSpell updates Spell
func (d *GormDatabase) UpdateSpell(spell *model.Spell) error {
	return d.DB.Save(spell).Error
}

// DeleteSpell deletes Spell
func (d *GormDatabase) DeleteSpell(id uint) error {
	return d.DB.Where("id = ?", id).
		Delete(&model.Spell{}).Error
}
