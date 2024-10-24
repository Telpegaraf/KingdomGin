package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"kingdom/mode"
	"kingdom/model"
	"kingdom/test"
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

	s.db = testdb.NewDBWithDefaultUser(s.T())

	s.a = &UserApi{DB: s.db}
}

func (s *UserSuite) Test_GetUsers() {
	first := s.db.NewUser(2)
	second := s.db.NewUser(5)

	s.a.GetUsers(s.ctx)

	assert.Equal(s.T(), 200, s.recorder.Code)
	test.BodyEquals(s.T(), []*model.UserExternal{externalOf(first), externalOf(second)}, s.recorder)
}

func externalOf(user *model.User) *model.UserExternal {
	return &model.UserExternal{Username: user.Username, Admin: user.Admin, ID: user.ID}
}
