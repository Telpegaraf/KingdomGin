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
	Domains         []Domain `gorm:"many2many:god_domains"`
}
