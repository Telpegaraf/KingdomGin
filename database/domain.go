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

// GetDomains returns all Domains
func (d *GormDatabase) GetDomains() ([]*model.Domain, error) {
	var domains []*model.Domain
	err := d.DB.Find(&domains).Error
	return domains, err
}

// UpdateDomain updates Domain
func (d *GormDatabase) UpdateDomain(domain *model.Domain) error { return d.DB.Save(domain).Error }

// DeleteDomain deletes Domain
func (d *GormDatabase) DeleteDomain(id uint) error {
	return d.DB.Where("id = ?", id).Delete(&model.Domain{}).Error
}
