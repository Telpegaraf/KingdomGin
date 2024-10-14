package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kingdom/model"
	"log"
)

func (s *DatabaseSuite) TestUser() {
	user, err := s.db.GetUserByID(55)
	require.NoError(s.T(), err)
	assert.Nil(s.T(), user, "not existing user")

	user, err = s.db.GetUserByUsername("testuser")
	require.NoError(s.T(), err)
	assert.Nil(s.T(), user, "not existing user")

	admin, err := s.db.GetUserByID(1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), admin, "on bootup the first user should be automatically created")

	adminCount, err := s.db.CountUser("admin = ?", true)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 1, adminCount, 1, "there is initially one admin")

	users, err := s.db.GetUsers()
	require.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	assert.Contains(s.T(), users, admin)

	testUser := &model.User{Username: "testuser", Password: []byte{1, 2, 3, 4}}
	s.db.CreateUser(testUser)
	assert.NotEqual(s.T(), 0, testUser.ID, "on create user a new id should be assigned")
	userCount, err := s.db.CountUser()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 2, userCount, "two users should exist")

	user, err = s.db.GetUserByUsername("testuser")
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testUser.Username, user.Username)
	assert.Equal(s.T(), testUser.Admin, user.Admin)
	assert.Equal(s.T(), testUser.Password, user.Password)
	assert.Equal(s.T(), testUser.Email, user.Email)

	users, err = s.db.GetUsers()
	require.NoError(s.T(), err)
	assert.Len(s.T(), users, 2)
	log.Println(testUser)
	assert.Contains(s.T(), users, admin)
	//assert.Contains(s.T(), users, testUser)
	// TODO To Solve Problem with Characters
	testUser.Username = "testuser2"
	require.NoError(s.T(), s.db.UpdateUser(testUser))

	testUser2, err := s.db.GetUserByUsername("testuser2")
	require.NoError(s.T(), err)
	assert.Equal(s.T(), testUser.Username, testUser2.Username)
	assert.Equal(s.T(), testUser.Password, testUser2.Password)

	users, err = s.db.GetUsers()
	require.NoError(s.T(), err)
	assert.Len(s.T(), users, 2)

	require.NoError(s.T(), s.db.DeleteUserByID(testUser2.ID))
	users, err = s.db.GetUsers()
	require.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	assert.Contains(s.T(), users, admin)

	s.db.DeleteUserByID(admin.ID)
	users, err = s.db.GetUsers()
	require.NoError(s.T(), err)
	assert.Empty(s.T(), users)
}
