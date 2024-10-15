package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
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
