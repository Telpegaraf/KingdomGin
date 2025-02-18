package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kingdom/auth/password"
	"kingdom/model"
	"kingdom/test"
	"testing"
)

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

type DatabaseSuite struct {
	suite.Suite
	db     *GormDatabase
	tmpDir test.TmpDir
}

func (s *DatabaseSuite) BeforeTest(suiteName, testName string) {
	s.tmpDir = test.NewTmpDir("kingdom_databasesuite")
	db, err := gorm.Open(sqlite.Open(s.tmpDir.Path("test.db")), &gorm.Config{})
	err = db.AutoMigrate(
		new(model.User),
		new(model.Tradition),
		new(model.Trait),
		new(model.Action),
		new(model.Skill),
		new(model.Race),
		new(model.Ancestry),
		new(model.Feat),
		new(model.Background),
		new(model.Item),
		new(model.Armor),
		new(model.Weapon),
		new(model.Gear),
		new(model.Character),
		new(model.Domain),
		new(model.God))
	if err != nil {
		return
	}
	userCount := int64(0)
	db.Find(new(model.User)).Count(&userCount)
	if userCount == 0 {
		db.Create(&model.User{
			Username: "admin",
			Password: password.CreatePassword("adminPassword", 10),
			Email:    "admin@example.com",
			Admin:    true})
	}
	assert.Nil(s.T(), err)
	s.db = &GormDatabase{DB: db}
}
