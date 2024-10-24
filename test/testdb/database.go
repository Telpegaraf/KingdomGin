package testdb

import (
	"fmt"
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

func NewDB(t *testing.T) *Database {
	db, err := gorm.Open(sqlite.Open("file:%s?mode=memory&cache=shared"), &gorm.Config{})
	db.AutoMigrate(new(model.User), new(model.Character), new(model.Domain), new(model.God))
	assert.Nil(t, err)
	assert.NotNil(t, db)
	tdb := &database.GormDatabase{
		DB: db,
	}
	return &Database{GormDatabase: tdb, t: t}
}

// NewUser creates a user and returns the user.
func (d *Database) NewUser(id uint) *model.User {
	return d.NewUserWithName(id, "user"+fmt.Sprint(id), "email"+fmt.Sprint(id)+"@example.com")
}

// NewUserWithName creates a user with a name and returns the user.
func (d *Database) NewUserWithName(id uint, name string, email string) *model.User {
	user := &model.User{
		ID:       id,
		Username: name,
		Email:    email,
		Password: password.CreatePassword(name, 10),
	}
	d.CreateUser(user)
	return user
}
