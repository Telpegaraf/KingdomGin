package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kingdom/auth/password"
	"kingdom/model"
	"os"
)

type GormDatabase struct {
	DB *gorm.DB
}

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

	sqlData, err := os.ReadFile("sqlSchema.sql")
	if err != nil {
		return nil, err
	}

	if err := db.Exec(string(sqlData)).Error; err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		new(model.User),
		new(model.Tradition),
		new(model.Skill),
		new(model.CharacterClass),
		new(model.ClassFeature),
		new(model.SkillFeature),
		new(model.Character),
		new(model.Domain),
		new(model.God),
		new(model.Action),
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
		new(model.Slot),
		new(model.CharacterBoost),
		new(model.Spell),
		new(model.CharacterDefence),
		new(model.CharacterSpell),
		new(model.CharacterFeat),
		new(model.CharacterSkill),
		new(model.CharacterInfo),
		new(model.UserCode),
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
