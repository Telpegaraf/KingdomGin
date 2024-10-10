package database

import "kingdom/model"

// CreateAncestry creates new Ancestry object
func (d *GormDatabase) CreateAncestry(ancestry *model.Ancestry) error {
	return d.DB.Create(ancestry).Error
}
