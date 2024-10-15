package model

type Item struct {
	ID          uint    `gorm:"primary_key"`
	Name        string  `gorm:"unique;type:varchar(127)"`
	Description string  `gorm:"type:text"`
	Bulk        float64 `gorm:"type:decimal(10,3);default:0.001"`
	Level       uint8   `gorm:"default:1;not null"`
	Price       string  `gorm:"type:varchar(127)"`
	OwnerID     uint
	OwnerType   string
}

type Armor struct {
	ID         uint `gorm:"primary_key;AUTO_INCREMENT"`
	ArmorClass uint
	Item       Item `gorm:"polymorphic:Owner;"`
}

type Weapon struct {
	ID           uint  `gorm:"primary_key;AUTO_INCREMENT"`
	DiceQuantity uint8 `gorm:"default:1;not null"`
	Dice         uint8 `gorm:"default:4;not null"`
	Damage       uint8 `gorm:"default:1;not null1"`
	Item         Item  `gorm:"polymorphic:Owner;"`
}

type Gear struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`
}

//type CreateItem struct {
//	Name        string  `json:"name" query:"name" form:"name" binding:"required"`
//	Description string  `json:"description" query:"description" form:"description" binding:"required"`
//	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk" binding:"required"`
//}
//
//type UpdateItem struct {
//	Name        string  `json:"name" query:"name" form:"name"`
//	Description string  `json:"description" query:"description" form:"description"`
//	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk"`
//}

//type Weapon struct {
//	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
//	OwnerID   uint
//	Damage    uint
//	OwnerType string
//}

//
//func (a *Armor) BeforeCreate(tx *gorm.DB) {
//	a.OwnerType = "items"
//	return
//}

//func (a *Weapon) BeforeUpdate(tx *gorm.DB) {
//	a.OwnerType = "items"
//	return
//}
