package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestFeat() {
	feats, err := s.db.GetFeats()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), feats)

	testFeat := &model.Feat{
		Name:        "Test Name",
		Description: "Test Description",
		Level:       1,
	}
	err = s.db.CreateFeat(testFeat)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testFeat.Level, uint8(1))

	backgrounds, err := s.db.GetBackgrounds()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), backgrounds)

	testBackground := &model.Background{
		Name:        "Test Background",
		Description: "Test Description",
		FeatID:      testFeat.ID,
	}
	require.NoError(s.T(), s.db.CreateBackground(testBackground))
	assert.Equal(s.T(), testBackground.Description, "Test Description")
	assert.Equal(s.T(), testBackground.FeatID, uint(1))

	backgrounds, err = s.db.GetBackgrounds()
	require.NoError(s.T(), err)
	assert.Len(s.T(), backgrounds, 1)

	testFeat.Name = "Test Feat Name"
	err = s.db.UpdateFeat(testFeat)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testFeat.Name, "Test Feat Name")

	testBackground.Name = "Test Background Name"
	err = s.db.UpdateBackground(testBackground)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testBackground.Name, "Test Background Name")

	err = s.db.DeleteFeat(testFeat.ID)
	require.NoError(s.T(), err)
	feats, err = s.db.GetFeats()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), feats)

	err = s.db.DeleteBackground(testBackground.ID)
	require.NoError(s.T(), err)
	backgrounds, err = s.db.GetBackgrounds()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), backgrounds)
}
