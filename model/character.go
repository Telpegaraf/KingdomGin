package model

type Character struct {
	ID               uint           `gorm:"primaryKey"`
	Name             string         `gorm:"not null;type:varchar(120)"`
	Alias            string         `gorm:"type:varchar(120)"`
	LastName         string         `gorm:"type:varchar(120)" json:"last_name"`
	Level            int8           `gorm:"default:1"`
	Race             Race           `gorm:"foreignKey:RaceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ancestry         Ancestry       `gorm:"foreignKey:AncestryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Background       Background     `gorm:"foreignKey:BackgroundID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterClass   CharacterClass `gorm:"foreignKey:CharacterClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID           uint
	RaceID           uint
	AncestryID       uint
	BackgroundID     uint
	CharacterClassID uint
	Attribute        Attribute        `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterSpell   []CharacterSpell `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterItem    []CharacterItem  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Slot             []Slot           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Boost            CharacterBoost   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterDefence CharacterDefence `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CharacterFeat    CharacterFeat    `gorm:"constraint:OnUpdate:CASCADE,onDelete:CASCADE;"`
	CharacterSkill   []CharacterSkill `gorm:"constraint:OnUpdate:CASCADE,onDelete:CASCADE;"`
	CharacterInfo    CharacterInfo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCharacter struct {
	Name             string `binding:"required" json:"name" query:"name" form:"name"`
	Alias            string `json:"alias" query:"alias" form:"alias"`
	LastName         string `json:"last_name" query:"last_name" form:"last_name"`
	CharacterClassID uint   `json:"character_class_id" query:"character_class_id" form:"character_class_id"`
	RaceID           uint   `json:"race_id" query:"race_id" form:"race_id"`
	AncestryID       uint   `json:"ancestry_id" query:"ancestry_id" form:"ancestry_id" binding:"required"`
	BackgroundID     uint   `json:"background_id" query:"background_id" form:"background_id" binding:"required"`
}

type CharacterUpdate struct {
	Name     string `json:"name" query:"name" form:"name"`
	Alias    string `json:"alias" query:"alias" form:"alias"`
	LastName string `json:"last_name" query:"last_name" form:"last_name"`
	Level    int8   `json:"level" query:"level" form:"level"`
}

type CharacterExternal struct {
	ID                  uint             `json:"id"`
	Name                string           `binding:"required" json:"name" query:"name" form:"name"`
	Alias               string           `json:"alias" query:"alias" form:"alias"`
	LastName            string           `json:"last_name" query:"last_name" form:"last_name"`
	Level               int8             `json:"level" query:"level" form:"level"`
	UserID              uint             `json:"user_id" query:"user_id" form:"user_id"`
	Attribute           Attribute        `json:"attribute" query:"attribute" form:"attribute"`
	CharacterItem       []CharacterItem  `json:"character_item" query:"character_item" form:"character_item"`
	Slot                []Slot           `json:"slot" query:"slot" form:"slot"`
	CharacterBoost      CharacterBoost   `json:"character_boost" query:"character_boost" form:"character_boost"`
	CharacterClassID    uint             `json:"character_class_id" query:"character_class_id" form:"character_class_id"`
	CharacterClass      CharacterClass   `json:"character_class" query:"character_class" form:"character_class"`
	CharacterRace       Race             `json:"character_race" query:"character_race" form:"character_race"`
	CharacterSkill      []CharacterSkill `json:"character_skill" query:"character_skill" form:"character_skill"`
	CharacterInfo       CharacterInfo    `json:"character_info"`
	CharacterAncestry   Ancestry         `json:"character_ancestry" query:"character_ancestry" form:"character_ancestry"`
	CharacterBackground Background       `json:"character_background"`
	CharacterDefence    CharacterDefence `json:"character_defence"`
}
