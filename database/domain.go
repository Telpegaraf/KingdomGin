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
		return domain, err
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
	return d.DB.Where("id = ?", id).
		Delete(&model.Domain{}).Error
}

// FindDomains Return Domain's ID in God object
func (d *GormDatabase) FindDomains(domainIDs []model.DomainID) ([]model.Domain, error) {
	var domains []model.Domain
	var ids []uint
	for _, domainID := range domainIDs {
		ids = append(ids, domainID.ID)
	}
	err := d.DB.Where("id IN (?)", ids).Find(&domains).Error
	return domains, err
}
