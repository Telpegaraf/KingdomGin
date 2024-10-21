package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestTrait() {
	trait, err := s.db.GetTraitByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), trait, "trait should be absent")

	testTrait := &model.Trait{Name: "Trait test", Description: "Test Description"}
	err = s.db.CreateTrait(testTrait)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Trait test", testTrait.Name)

	traits, err := s.db.GetTraits()
	require.NoError(s.T(), err)
	assert.Len(s.T(), traits, 1)
	assert.Contains(s.T(), traits, testTrait)

	testTrait.Name = "Trait 2"
	require.NoError(s.T(), s.db.UpdateTrait(testTrait))
	testTrait.Description = "Test Description 2"
	require.NoError(s.T(), s.db.UpdateTrait(testTrait))
	testTrait2, err := s.db.GetTraitByID(2)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), testTrait2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Trait 2", testTrait.Name)
	assert.Equal(s.T(), "Test Description 2", testTrait.Description)

	err = s.db.DeleteTrait(testTrait.ID)
	require.NoError(s.T(), err)
	traits, err = s.db.GetTraits()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), traits)
}
