package model

type Slot struct {
	ID             uint `gorm:"primary_key;AUTO_INCREMENT"`
	CharacterID    uint `gorm:"unique;"`
	ArmorID        *uint
	FirstWeaponID  *uint
	SecondWeaponID *uint

	Character    Character     `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Armor        CharacterItem `gorm:"foreignKey:ArmorID;references:ID;constraint:OnUpdate:CASCADE"`
	FirstWeapon  CharacterItem `gorm:"foreignKey:FirstWeaponID;references:ID;constraint:OnUpdate:CASCADE"`
	SecondWeapon CharacterItem `gorm:"foreignKey:SecondWeaponID;references:ID;constraint:OnUpdate:CASCADE"`
}

type SlotUpdate struct {
	CharacterID    uint  `json:"character_id" query:"character_id" form:"character_id"`
	ArmorID        *uint `json:"armor_id" query:"armor_id" form:"armor_id"`
	FirstWeaponID  *uint `json:"first_weapon_id" query:"first_weapon_id" form:"first_weapon_id"`
	SecondWeaponID *uint `json:"second_weapon_id" query:"second_weapon_id" form:"second_weapon_id"`
}

type SlotExternal struct {
	ID             uint  `json:"id" query:"id" form:"id"`
	CharacterID    uint  `json:"character_id" query:"character_id" form:"character_id"`
	ArmorID        *uint `json:"armor_id" query:"armor_id" form:"armor_id"`
	FirstWeaponID  *uint `json:"first_weapon_id" query:"first_weapon_id" form:"first_weapon_id"`
	SecondWeaponID *uint `json:"second_weapon_id" query:"second_weapon_id" form:"second_weapon_id"`
}
