package database

import "kingdom/model"

// CreateSpell creates new Spell object
func (d *GormDatabase) CreateSpell(spell *model.Spell) error {
	return d.DB.Create(spell).Error
}
