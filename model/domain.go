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
