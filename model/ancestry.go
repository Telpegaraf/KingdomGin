package model

type Ancestry struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `gorm:"unique"`
	Description string `gorm:"type:text"`
	RaceID      uint
}

type AncestryCreate struct {
	Name        string `json:"name" binding:"required" query:"name" form:"name"`
	Description string `json:"description" binding:"required" query:"description" form:"description"`
	RaceID      uint   `json:"race_id" binding:"required" query:"race_id" form:"race_id"`
}

type AncestryUpdate struct {
	Name        string `json:"name" query:"name" form:"name"`
	Description string `json:"description" query:"description" form:"description"`
	RaceID      uint   `json:"race_id" query:"race_id" form:"race_id"`
}

type AncestryExternal struct {
	ID          uint   `json:"id" query:"id" form:"id"`
	Name        string `json:"name" query:"name" form:"name"`
	Description string `json:"description" query:"description" form:"description"`
	RaceID      uint   `json:"race_id" query:"race_id" form:"race_id"`
}
