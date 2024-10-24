package model

type CharacterInfo struct {
	ID          uint    `gorm:"primary_key;AUTO_INCREMENT"`
	ClassDC     uint8   `gorm:"default:13"`
	HeroPoint   uint8   `gorm:"default:1"`
	MaxBulk     float64 `gorm:"type:decimal(10,3)"`
	Bulk        float64 `gorm:"type:decimal(10,3);default:0"`
	CharacterID uint    `gorm:"unique;"`
}
