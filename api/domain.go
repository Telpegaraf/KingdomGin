package api

import "kingdom/model"

type DomainDatabase interface {
	GetDomainByID(id uint) (*model.Domain, error)
	CreateDomain(domain *model.Domain) error
}

type DomainApi struct {
	DB DomainDatabase
}
