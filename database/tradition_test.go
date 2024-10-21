package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestTradition() {
	tradition, err := s.db.GetTraditionByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), tradition, "not existing user")

	testTradition := &model.Tradition{Name: "Tradition test", Description: "Test Description"}
	err = s.db.CreateTradition(testTradition)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Tradition test", testTradition.Name)

	traditions, err := s.db.GetTraditions()
	require.NoError(s.T(), err)
	assert.Len(s.T(), traditions, 1)
	assert.Contains(s.T(), traditions, testTradition)

	testTradition.Name = "Tradition 2"
	require.NoError(s.T(), s.db.UpdateTradition(testTradition))
	testTradition.Description = "Test Description 2"
	require.NoError(s.T(), s.db.UpdateTradition(testTradition))
	testTradition2, err := s.db.GetTraditionByID(2)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), testTradition2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Tradition 2", testTradition.Name)
	assert.Equal(s.T(), "Test Description 2", testTradition.Description)

	err = s.db.DeleteTradition(testTradition.ID)
	require.NoError(s.T(), err)
	traditions, err = s.db.GetTraditions()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), traditions)
}
