package model

type God struct {
	ID              uint     `gorm:"primary_key;AUTO_INCREMENT"`
	Name            string   `gorm:"unique;not null;type:varchar(120)"`
	Alias           string   `gorm:"unique;type:varchar(120)"`
	Edict           string   `gorm:"type:varchar(255)"`
	Anathema        string   `gorm:"type:varchar(255)"`
	AreasOfInterest string   `gorm:"type:varchar(255)"`
	Temples         string   `gorm:"type:varchar(255)"`
	Worships        string   `gorm:"type:varchar(255)"`
	SacredAnimals   string   `gorm:"type:varchar(255)"`
	SacredColors    string   `gorm:"type:varchar(255)"`
	ChosenWeapon    string   `gorm:"type:varchar(255)"`
	Alignment       string   `gorm:"type:varchar(255)"`
	Description     string   `gorm:"type:text"`
	Domains         []Domain `gorm:"many2many:god_domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type GodCreate struct {
	Name            string     `json:"name" binding:"required" query:"name" form:"name"`
	Alias           string     `json:"alias" binding:"required" query:"alias" form:"alias"`
	Edict           string     `json:"edict" binding:"required" query:"edict" form:"edict"`
	Anathema        string     `json:"anathema" binding:"required" query:"anathema" form:"anathema"`
	AreasOfInterest string     `json:"areas_of_interest" query:"areas_of_interest" form:"areas_of_interest"`
	Temples         string     `json:"temples" binding:"required" query:"temples" form:"temples"`
	Worships        string     `json:"worships" query:"worships" form:"worships"`
	SacredAnimals   string     `json:"sacred_animals" query:"sacred_animals" form:"sacred_animals"`
	SacredColors    string     `json:"sacred_colors" query:"sacred_colors" form:"sacred_colors"`
	ChosenWeapon    string     `json:"chosen_weapon" query:"chosen_weapon" form:"chosen_weapon"`
	Alignment       string     `json:"alignment" query:"alignment" form:"alignment"`
	Description     string     `json:"description" query:"description" form:"description"`
	Domains         []DomainID `json:"domains" binding:"required" query:"domains"`
}

type GodUpdate struct {
	Name            string     `json:"name" binding:"required" query:"name" form:"name"`
	Alias           string     `json:"alias" binding:"required" query:"alias" form:"alias"`
	Edict           string     `json:"edict" binding:"required" query:"edict" form:"edict"`
	Anathema        string     `json:"anathema" binding:"required" query:"anathema" form:"anathema"`
	AreasOfInterest string     `json:"areas_of_interest" query:"areas_of_interest" form:"areas_of_interest"`
	Temples         string     `json:"temples" binding:"required" query:"temples" form:"temples"`
	Worships        string     `json:"worships" query:"worships" form:"worships"`
	SacredAnimals   string     `json:"sacred_animals" query:"sacred_animals" form:"sacred_animals"`
	SacredColors    string     `json:"sacred_colors" query:"sacred_colors" form:"sacred_colors"`
	ChosenWeapon    string     `json:"chosen_weapon" query:"chosen_weapon" form:"chosen_weapon"`
	Alignment       string     `json:"alignment" query:"alignment" form:"alignment"`
	Description     string     `json:"description" query:"description" form:"description"`
	Domains         []DomainID `json:"domains" binding:"required" query:"domains"`
}
