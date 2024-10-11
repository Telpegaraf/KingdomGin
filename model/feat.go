package model

type Feat struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `gorm:"unique;not null"`
	Description string
	Level       int8 `gorm:"default:1;not null"`
	Background  []Background
}

type CreateFeat struct {
	Name        string `json:"name" binding:"required" query:"name"`
	Description string `json:"description" query:"description"`
}
