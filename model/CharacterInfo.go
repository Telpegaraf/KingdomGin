package model

type CharacterInfo struct {
	iD        uint `gorm:"primary_key;AUTO_INCREMENT"`
	HeroPoint uint8
}
