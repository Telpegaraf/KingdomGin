package model

type CharacterInfo struct {
	iD        uint  `gorm:"primary_key;AUTO_INCREMENT"`
	ClassDC   uint8 `gorm:"default:10"`
	HeroPoint uint8
}
