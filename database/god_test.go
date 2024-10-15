package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestGod() {
	emptyGod, err := s.db.GetDomainByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), emptyGod, "God object must be absent")

	domain := &model.Domain{
		Name:        "First Domain",
		Description: "First Description",
	}

	err = s.db.CreateDomain(domain)
	require.NoError(s.T(), err)
	firstDomain, err := s.db.GetDomainByID(1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), firstDomain)

	firstGod := &model.God{
		Name:            "Test God",
		Description:     "Test Description",
		Alias:           "Test Alias",
		Edict:           "Test Edict",
		Anathema:        "Test Anathema",
		AreasOfInterest: "Test AreasOfInterest",
		Temples:         "Test Temples",
		Worships:        "Test Worships",
		SacredColors:    "Test SacredColors",
		SacredAnimals:   "Test SacredAnimals",
		ChosenWeapon:    "Test ChosenWeapons",
		Alignment:       "Test Alignment",
		Domains:         []model.Domain{*domain},
	}
	err = s.db.CreateGod(firstGod)
	require.NoError(s.T(), err)
	god, err := s.db.GetGodByID(1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), god)
	assert.Equal(s.T(), "Test God", god.Name)

	firstGod.Name = "New Test God"
	require.NoError(s.T(), s.db.UpdateGod(firstGod))
	secondGod, err := s.db.GetGodByID(1)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), secondGod.Name, "New Test God")
	assert.Equal(s.T(), secondGod.Description, "Test Description")

	require.NoError(s.T(), s.db.DeleteGod(1))
	gods, err := s.db.GetGods()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), gods)
}
