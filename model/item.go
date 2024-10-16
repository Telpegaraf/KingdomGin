package model

type Item struct {
	ID            uint            `gorm:"primary_key"`
	Name          string          `gorm:"unique;type:varchar(127)"`
	Description   string          `gorm:"type:text"`
	Bulk          float64         `gorm:"type:decimal(10,3);default:0.001"`
	Level         uint8           `gorm:"default:1;not null"`
	Price         string          `gorm:"type:varchar(127)"`
	OwnerID       uint            `gorm:"uniqueIndex:idx_owner_id_owner_type"`
	OwnerType     string          `gorm:"uniqueIndex:idx_owner_id_owner_type"`
	CharacterItem []CharacterItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ItemExternal struct {
	ID          uint    `json:"id" query:"id" form:"id"`
	Name        string  `json:"name" query:"name" binding:"required" form:"name"`
	Description string  `json:"description" query:"description" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk"`
	Level       uint8   `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" binding:"required" form:"price"`
	OwnerID     uint    `json:"owner_id" query:"owner_id" form:"owner_id"`
	OwnerType   string  `json:"owner_type" query:"owner_type" form:"owner_type"`
}

type Armor struct {
	ID         uint `gorm:"primary_key;AUTO_INCREMENT"`
	ArmorClass uint8
	Item       Item `gorm:"polymorphic:Owner;"`
}

type Weapon struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	DiceQuantity uint8  `gorm:"default:1;not null"`
	Dice         uint8  `gorm:"default:4;not null"`
	Damage       uint8  `gorm:"default:1;not null"`
	Item         Item   `gorm:"polymorphic:Owner;"`
	DamageType   string `gorm:"type:varchar(127);not null"`
}

type Gear struct {
	ID   uint `gorm:"primary_key;AUTO_INCREMENT"`
	Item Item `gorm:"polymorphic:Owner;"`
}

type CreateArmor struct {
	Name        string  `json:"name" query:"name" binding:"required" form:"name"`
	Description string  `json:"description" query:"description" binding:"required" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk"`
	Level       *uint8  `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" binding:"required" form:"price"`
	ArmorClass  *uint8  `json:"armor_class" query:"armor_class" form:"armor_class"`
}

type UpdateArmor struct {
	Name        string  `json:"name" query:"name" form:"name"`
	Description string  `json:"description" query:"description" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk"`
	Level       *uint8  `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" form:"price"`
	ArmorClass  *uint8  `json:"armor_class" query:"armor_class" form:"armor_class"`
}

type ArmorExternal struct {
	ID          uint    `json:"id" query:"id" form:"id"`
	Name        string  `json:"name" query:"name" binding:"required" form:"name"`
	Description string  `json:"description" query:"description" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk"`
	Level       uint8   `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" binding:"required" form:"price"`
	ArmorClass  uint8   `json:"armor_class" query:"armor_class" form:"armor_class"`
	ItemID      uint    `json:"item_id" query:"item_id" form:"item_id"`
}

type CreateWeapon struct {
	Name         string  `json:"name" query:"name" binding:"required" form:"name"`
	Description  string  `json:"description" query:"description" binding:"required" form:"description"`
	Bulk         float64 `json:"bulk" query:"bulk" binding:"required" form:"bulk" example:"0.1"`
	Level        *uint8  `json:"level" query:"level" form:"level"`
	Price        string  `json:"price" query:"price" binding:"required" form:"price"`
	DiceQuantity uint8   `json:"diceQuantity" query:"dice_quantity" form:"dice_quantity" binding:"required" example:"1"`
	Dice         uint8   `json:"dice" query:"dice" form:"dice" binding:"required" example:"4"`
	Damage       *uint8  `json:"damage" query:"damage" form:"damage"`
	DamageType   string  `json:"damage_type" query:"damage_type" form:"damage_type" binding:"required"`
}

type UpdateWeapon struct {
	Name         string  `json:"name" query:"name" form:"name"`
	Description  string  `json:"description" query:"description" form:"description"`
	Bulk         float64 `json:"bulk" query:"bulk" form:"bulk"`
	Level        *uint8  `json:"level" query:"level" form:"level"`
	Price        string  `json:"price" query:"price" form:"price"`
	DiceQuantity uint8   `json:"diceQuantity" query:"dice_quantity" form:"dice_quantity"`
	Dice         uint8   `json:"dice" query:"dice" form:"dice"`
	Damage       *uint8  `json:"damage" query:"damage" form:"damage"`
	DamageType   string  `json:"damage_type" query:"damage_type" form:"damage_type"`
}

type WeaponExternal struct {
	ID           uint    `json:"id" query:"id" form:"id"`
	Name         string  `json:"name" query:"name" binding:"required" form:"name"`
	Description  string  `json:"description" query:"description" form:"description"`
	Bulk         float64 `json:"bulk" query:"bulk" form:"bulk"`
	Level        uint8   `json:"level" query:"level" form:"level" example:"1"`
	Price        string  `json:"price" query:"price" binding:"required" form:"price"`
	DiceQuantity uint8   `json:"diceQuantity" query:"dice_quantity" form:"dice_quantity" example:"1"`
	Dice         uint8   `json:"dice" query:"dice" form:"dice" example:"1"`
	Damage       uint8   `json:"damage" query:"damage" form:"damage" example:"1"`
	DamageType   string  `json:"damage_type" query:"damage_type" form:"damage_type"`
	ItemID       uint    `json:"item_id" query:"item_id" form:"item_id"`
}

type CreateGear struct {
	Name        string  `json:"name" query:"name" binding:"required" form:"name"`
	Description string  `json:"description" query:"description" binding:"required" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" binding:"required" form:"bulk" example:"1"`
	Level       *uint8  `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" binding:"required" form:"price"`
}

type UpdateGear struct {
	Name        string  `json:"name" query:"name" form:"name"`
	Description string  `json:"description" query:"description" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk" example:"1"`
	Level       *uint8  `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" form:"price"`
}

type GearExternal struct {
	ID          uint    `json:"id" query:"id" form:"id"`
	Name        string  `json:"name" query:"name" form:"name"`
	Description string  `json:"description" query:"description" form:"description"`
	Bulk        float64 `json:"bulk" query:"bulk" form:"bulk" example:"1"`
	Level       uint8   `json:"level" query:"level" form:"level"`
	Price       string  `json:"price" query:"price" form:"price"`
	ItemID      uint    `json:"item_id" query:"item_id" form:"item_id"`
}
