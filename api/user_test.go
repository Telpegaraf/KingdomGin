package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"kingdom/mode"
	"kingdom/test/testdb"
	"net/http/httptest"
	"testing"
)

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

type UserSuite struct {
	suite.Suite
	db       *testdb.Database
	a        *UserApi
	ctx      *gin.Context
	recorder *httptest.ResponseRecorder
}

func (s *UserSuite) BeforeTest(suiteName, testName string) {
	mode.Set(mode.TestDev)
	s.recorder = httptest.NewRecorder()
	s.ctx, _ = gin.CreateTestContext(s.recorder)

	s.db = testdb.NewDB(s.T())

	s.a = &UserApi{DB: s.db}
}

//TODO fix problem with nil
//func (s *UserSuite) Test_GetUsers() {
//	first := s.db.NewUser(2)
//	second := s.db.NewUser(5)
//
//	s.a.GetUsers(s.ctx)
//
//	assert.Equal(s.T(), 200, s.recorder.Code)
//	test.BodyEquals(s.T(), []*model.UserExternal{externalOf(first), externalOf(second)}, s.recorder)
//}

//func (s *UserSuite) Test_GetCurrentUser() {
//	user := s.db.NewUser(5)
//
//	test.WithUser(s.ctx, 5)
//	s.a.GetCurrentUser(s.ctx)
//
//	assert.Equal(s.T(), 200, s.recorder.Code)
//	test.BodyEquals(s.T(), externalOf(user), s.recorder)
//}
//
//func externalOf(user *model.User) *model.UserExternal {
//	return &model.UserExternal{Username: user.Username, Admin: user.Admin, ID: user.ID}
//}
