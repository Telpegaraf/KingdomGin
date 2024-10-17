package model

type God struct {
	ID              uint     `gorm:"primary_key;AUTO_INCREMENT"`
	Name            string   `gorm:"unique;not null;type:varchar(120)"`
	Alias           string   `gorm:"unique;type:varchar(120); not null"`
	Edict           string   `gorm:"type:varchar(255); not null"`
	Anathema        string   `gorm:"type:varchar(255); not null"`
	AreasOfInterest string   `gorm:"type:varchar(255); not null"`
	Temples         string   `gorm:"type:varchar(255); not null"`
	Worships        string   `gorm:"type:varchar(255); not null"`
	SacredAnimals   string   `gorm:"type:varchar(255); not null"`
	SacredColors    string   `gorm:"type:varchar(255); not null"`
	ChosenWeapon    string   `gorm:"type:varchar(255); not null"`
	Alignment       string   `gorm:"type:varchar(255); not null"`
	Description     string   `gorm:"type:text; not null"`
	Domains         []Domain `gorm:"many2many:god_domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}

type GodCreate struct {
	Name            string     `json:"name" binding:"required" query:"name" form:"name" example:"Desna"`
	Alias           string     `json:"alias" binding:"required" query:"alias" form:"alias" example:"The Song of the Spheres"`
	Edict           string     `json:"edict" binding:"required" query:"edict" form:"edict"`
	Anathema        string     `json:"anathema" binding:"required" query:"anathema" form:"anathema"`
	AreasOfInterest string     `json:"areas_of_interest" binding:"required" query:"areas_of_interest" form:"areas_of_interest"`
	Temples         string     `json:"temples" binding:"required" query:"temples" form:"temples"`
	Worships        string     `json:"worships" binding:"required" query:"worships" form:"worships"`
	SacredAnimals   string     `json:"sacred_animals" binding:"required" query:"sacred_animals" form:"sacred_animals"`
	SacredColors    string     `json:"sacred_colors" binding:"required" query:"sacred_colors" form:"sacred_colors"`
	ChosenWeapon    string     `json:"chosen_weapon" binding:"required" query:"chosen_weapon" form:"chosen_weapon"`
	Alignment       string     `json:"alignment" binding:"required" query:"alignment" form:"alignment" example:"CG"`
	Description     string     `json:"description" query:"description" form:"description"`
	Domains         []DomainID `json:"domains" binding:"required" query:"domains"`
}

type GodUpdate struct {
	Name            string     `json:"name" query:"name" form:"name"`
	Alias           string     `json:"alias" query:"alias" form:"alias"`
	Edict           string     `json:"edict" query:"edict" form:"edict"`
	Anathema        string     `json:"anathema" query:"anathema" form:"anathema"`
	AreasOfInterest string     `json:"areas_of_interest" query:"areas_of_interest" form:"areas_of_interest"`
	Temples         string     `json:"temples" query:"temples" form:"temples"`
	Worships        string     `json:"worships" query:"worships" form:"worships"`
	SacredAnimals   string     `json:"sacred_animals" query:"sacred_animals" form:"sacred_animals"`
	SacredColors    string     `json:"sacred_colors" query:"sacred_colors" form:"sacred_colors"`
	ChosenWeapon    string     `json:"chosen_weapon" query:"chosen_weapon" form:"chosen_weapon"`
	Alignment       string     `json:"alignment" query:"alignment" form:"alignment"`
	Description     string     `json:"description" query:"description" form:"description"`
	Domains         []DomainID `json:"domains" query:"domains"`
}

type GodExternal struct {
	ID              uint     `json:"id" query:"id" form:"id"`
	Name            string   `json:"name" query:"name" form:"name"`
	Alias           string   `json:"alias" query:"alias" form:"alias"`
	Edict           string   `json:"edict" query:"edict" form:"edict"`
	Anathema        string   `json:"anathema" query:"anathema" form:"anathema"`
	AreasOfInterest string   `json:"areas_of_interest" query:"areas_of_interest" form:"areas_of_interest"`
	Temples         string   `json:"temples" query:"temples" form:"temples"`
	Worships        string   `json:"worships" query:"worships" form:"worships"`
	SacredAnimals   string   `json:"sacred_animals" query:"sacred_animals" form:"sacred_animals"`
	SacredColors    string   `json:"sacred_colors" query:"sacred_colors" form:"sacred_colors"`
	ChosenWeapon    string   `json:"chosen_weapon" query:"chosen_weapon" form:"chosen_weapon"`
	Alignment       string   `json:"alignment" query:"alignment" form:"alignment"`
	Description     string   `json:"description" query:"description" form:"description"`
	Domains         []Domain `json:"domains" query:"domains"`
}
