package database

import (
	"errors"
	"gorm.io/gorm"
	"kingdom/model"
)

// GetDomainByID returns Domain by ID
func (d *GormDatabase) GetDomainByID(id uint) (*model.Domain, error) {
	domain := new(model.Domain)
	err := d.DB.Find(domain, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if domain.ID == id {
		return domain, nil
	}
	return nil, err
}

// CreateDomain create new Domain
func (d *GormDatabase) CreateDomain(domain *model.Domain) error { return d.DB.Create(domain).Error }
