package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kingdom/auth/password"
	"kingdom/model"
)

func New(
	dsn,
	defaultUser string,
	defaultPass string,
	defaultEmail string,
	strength int,
	createDefaultUserIfNotExist bool) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		new(model.User),
		new(model.Character),
		new(model.Domain),
		new(model.God),
		new(model.CharacterClass),
		new(model.Attribute),
		new(model.Item),
		new(model.Feat),
		new(model.Race),
		new(model.Ancestry),
		new(model.Background),
		new(model.Item),
		new(model.CharacterItem),
		new(model.Armor),
		new(model.Weapon),
		new(model.Gear),
	); err != nil {
		return nil, err
	}

	userCount := int64(0)
	db.Find(new(model.User)).Count(&userCount)
	if createDefaultUserIfNotExist && userCount == 0 {
		db.Create(&model.User{
			Username: defaultUser,
			Password: password.CreatePassword(defaultPass, strength),
			Email:    defaultEmail,
			Admin:    true})
	}

	return &GormDatabase{DB: db}, nil
}

type GormDatabase struct {
	DB *gorm.DB
}
