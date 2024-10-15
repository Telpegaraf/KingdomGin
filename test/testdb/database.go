package testdb

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kingdom/auth/password"
	"kingdom/database"
	"kingdom/model"
	"testing"
)

type Database struct {
	*database.GormDatabase
	t *testing.T
}

func NewDBWithDefaultUser(t *testing.T) *Database {
	db, err := gorm.Open(sqlite.Open("file:%s?mode=memory&cache=shared"), &gorm.Config{})
	db.AutoMigrate(new(model.User), new(model.Character), new(model.Domain), new(model.God))
	userCount := int64(0)
	db.Find(new(model.User)).Count(&userCount)
	if userCount == 0 {
		db.Create(&model.User{
			Username: "admin",
			Password: password.CreatePassword("adminPassword", 10),
			Email:    "admin@example.com",
			Admin:    true})
	}
	assert.Nil(t, err)
	assert.NotNil(t, db)
	tdb := &database.GormDatabase{
		DB: db,
	}
	return &Database{GormDatabase: tdb, t: t}
}
