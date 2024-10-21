package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
)

func (s *DatabaseSuite) TestAction() {
	action, err := s.db.GetActionByID(1)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), action, "Action should be absent")

	testAction := &model.Action{Name: "Action test"}
	err = s.db.CreateAction(testAction)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Action test", testAction.Name)

	actions, err := s.db.GetActions()
	require.NoError(s.T(), err)
	assert.Len(s.T(), actions, 1)
	assert.Contains(s.T(), actions, testAction)

	testAction.Name = "Action 2"
	require.NoError(s.T(), s.db.UpdateAction(testAction))
	testAction2, err := s.db.GetActionByID(2)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), testAction2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "Action 2", testAction.Name)

	err = s.db.DeleteAction(testAction.ID)
	require.NoError(s.T(), err)
	actions, err = s.db.GetActions()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), actions)
}
