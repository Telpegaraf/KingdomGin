package model

type Domain struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique;not null;type:varchar(120)"`
	Description string `gorm:"type:text"`
	Gods        []God  `gorm:"many2many:god_domains;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DomainID struct {
	ID uint
}

type CreateDomain struct {
	Name        string `json:"name" query:"name" form:"name" binding:"required"`
	Description string `json:"description" query:"description" form:"description" binding:"required"`
}

type UpdateDomain struct {
	Name        string `json:"name" query:"name" form:"name" binding:"required"`
	Description string `json:"description" query:"description" form:"description" binding:"required"`
}
