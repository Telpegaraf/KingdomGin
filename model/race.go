package model

type Race struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"unique;not null"`
	Description  string `gorm:"not null"`
	HitPoints    int8   `gorm:"not null"`
	Size         string `gorm:"type:varchar(31);not null"`
	Speed        int8   `gorm:"not null"`
	AbilityBoost string `gorm:"not null"`
	Language     string
	Ancestry     []Ancestry
}

type RaceCreate struct {
	Name         string `json:"name" binding:"required" query:"name" form:"name"`
	Description  string `json:"description" binding:"required" query:"description" form:"description"`
	HitPoints    int8   `json:"hit_points" binding:"required" query:"hit_points" form:"hit_points"`
	Size         string `json:"size" binding:"required" query:"size" form:"size"`
	Speed        int8   `json:"speed" binding:"required" query:"speed" form:"speed"`
	AbilityBoost string `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	Language     string `json:"language" query:"language" form:"language"`
}

type RaceUpdate struct {
	Name         string `json:"name" query:"name" form:"name"`
	Description  string `json:"description" query:"description" form:"description"`
	HitPoints    int8   `json:"hit_points" query:"hit_points" form:"hit_points"`
	Size         string `json:"size" query:"size" form:"size"`
	Speed        int8   `json:"speed" query:"speed" form:"speed"`
	AbilityBoost string `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	Language     string `json:"language" query:"language" form:"language"`
}

type RaceExternal struct {
	ID           uint   `json:"id" query:"id" form:"id"`
	Name         string `json:"name" query:"name" form:"name"`
	Description  string `json:"description" query:"description" form:"description"`
	HitPoints    int8   `json:"hit_points" query:"hit_points" form:"hit_points"`
	Size         string `json:"size" query:"size" form:"size"`
	Speed        int8   `json:"speed" query:"speed" form:"speed"`
	AbilityBoost string `json:"ability_boost" query:"ability_boost" form:"ability_boost"`
	Language     string `json:"language" query:"language" form:"language"`
}
