package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"kingdom/model"
)

func (s *DatabaseSuite) TestDomain() {
	domain, err := s.db.GetDomainByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), domain)

	testDomain := &model.Domain{
		Name:        "Test Domain",
		Description: "Test Description",
	}
	err = s.db.CreateDomain(testDomain)
	require.NoError(s.T(), err, "First domain should be created successfully")

	domain, err = s.db.GetDomainByID(1)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), uint(1), domain.ID, "on create domain a new id should be assigned")
	assert.Equal(s.T(), testDomain.Name, domain.Name, "on create domain a new name should be assigned")
	assert.Equal(s.T(), testDomain.Description, domain.Description, "on create domain a new description should be assigned")

	testDomainTwo := &model.Domain{
		Name:        "Test Domain", //Same name, must be unique
		Description: "Another Test Description ",
	}
	err = s.db.CreateDomain(testDomainTwo)
	assert.Error(s.T(), gorm.ErrDuplicatedKey)

	testDomainThree := &model.Domain{
		Name:        "Test Domain Two",
		Description: "Test Description",
	}
	err = s.db.CreateDomain(testDomainThree)
	require.NoError(s.T(), err)
	domainThree, err := s.db.GetDomainByID(2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Test Domain Two", domainThree.Name)

	testDomainThree.Name = "Test Domain 2"
	require.NoError(s.T(), s.db.UpdateDomain(testDomainThree))
	testDomainThree.Name = "Test Domain"
	err = s.db.UpdateDomain(testDomainTwo)
	assert.Error(s.T(), gorm.ErrDuplicatedKey)
	domainThree, err = s.db.GetDomainByID(2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Test Domain 2", domainThree.Name)

	s.db.DeleteDomain(1)
	s.db.DeleteDomain(2)
	domains, err := s.db.GetDomains()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 0, len(domains))
}
