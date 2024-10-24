package database

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestRace() {
	race, err := s.db.GetRaceByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), race)

	testRace := &model.Race{
		Name:         "Test Race",
		Description:  "Test Description",
		HitPoint:     6,
		Size:         model.Medium,
		Speed:        30,
		AbilityBoost: 2,
		Language:     "Test Language",
	}
	err = s.db.CreateRace(testRace)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testRace.Name, "Test Race")

	testRace2 := &model.Race{
		Name:         "Test Race2",
		Description:  "Test Description",
		HitPoint:     6,
		Size:         "medium",
		Speed:        30,
		AbilityBoost: 3,
		Language:     "Test Language",
	}
	err = s.db.CreateRace(testRace2)
	assert.Error(s.T(), errors.New("invalid Square Size vale"))

	races, err := s.db.GetRaces()
	require.NoError(s.T(), err)
	assert.Len(s.T(), races, 1)
	assert.Contains(s.T(), races, testRace)

	testRace.Name = "Test Race 2"
	testRace.Description = "Test Description 2"
	testRace.HitPoint = 6
	testRace.Size = model.Small
	testRace.Speed = 35
	testRace.AbilityBoost = 2
	err = s.db.UpdateRace(testRace)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testRace.Name, "Test Race 2")
	assert.Equal(s.T(), testRace.Description, "Test Description 2")
	assert.Equal(s.T(), testRace.HitPoint, uint16(6))
	assert.Equal(s.T(), testRace.Size, model.Small)
	assert.Equal(s.T(), testRace.Speed, uint8(35))
	assert.Equal(s.T(), testRace.AbilityBoost, uint8(2))
	assert.Equal(s.T(), testRace.Language, "Test Language")

	ancestries, err := s.db.GetAncestries()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), ancestries)

	testAncestry := &model.Ancestry{
		Name:        "Test Ancestry",
		Description: "Test Description",
		RaceID:      1,
	}
	err = s.db.CreateAncestry(testAncestry)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testAncestry.Name, "Test Ancestry")

	ancestries, err = s.db.GetAncestries()
	require.NoError(s.T(), err)
	assert.Len(s.T(), ancestries, 1)

	err = s.db.DeleteRaceByID(1)
	require.NoError(s.T(), err)
	races, err = s.db.GetRaces()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), races)

	err = s.db.DeleteAncestry(1)
	require.NoError(s.T(), err)
	ancestries, err = s.db.GetAncestries()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), ancestries)
}
