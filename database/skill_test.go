package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestSkill() {
	skill, err := s.db.GetSkillByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), skill)

	testSkill := &model.Skill{Name: "Test Skill", Description: "Test Description"}
	err = s.db.CreateSkill(testSkill)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Test Skill", testSkill.Name)

	skills, err := s.db.GetSkills()
	require.NoError(s.T(), err)
	assert.Len(s.T(), skills, 1)
	assert.Contains(s.T(), skills, testSkill)

	testSkill.Name = "Test Skill2"
	testSkill.Description = "Test Description2"
	err = s.db.UpdateSkill(testSkill)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Test Skill2", testSkill.Name)
	assert.Equal(s.T(), "Test Description2", testSkill.Description)

	err = s.db.DeleteSkill(testSkill.ID)
	require.NoError(s.T(), err)
	skills, err = s.db.GetSkills()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), skills)
}
