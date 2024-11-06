package model

type Spell struct {
	ID             uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string
	Description    string `gorm:"type:text"`
	Component      string `gorm:"type:text"`
	Range          string `gorm:"type:varchar(255)"`
	Area           string `gorm:"type:varchar(255)"`
	Duration       string `gorm:"type:varchar(255)"`
	Target         string `gorm:"type:varchar(255)"`
	Rank           uint8
	Ritual         bool             `gorm:"type:bool;default:false"`
	School         *School          `gorm:"type:school"`
	CharacterSpell []CharacterSpell `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cast           string
	Tradition      []Tradition `gorm:"many2many:spell_traditions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Traits         []Trait     `gorm:"many2many:spell_traits;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SpellCreate struct {
	Name        string `json:"name" query:"name"`
	Description string `json:"description" query:"description"`
	Component   string `json:"component" query:"component"`
	Range       string `json:"range" query:"range"`
	Area        string `json:"area" query:"area"`
	Duration    string `json:"duration" query:"duration"`
	Target      string `json:"target" query:"target"`
	Rank        uint8  `json:"rank" query:"rank"`
	Ritual      bool   `json:"ritual" query:"ritual"`
	School      School `json:"school" query:"school"`
	Cast        string `json:"cast" query:"cast"`
	TraitsID    []uint `json:"traits_id" query:"traits_id"`
	TraditionID []uint `json:"tradition_id" query:"tradition_id"`
}

type SpellUpdate struct {
	Name        string `json:"name" query:"name"`
	Description string `json:"description" query:"description"`
	Component   string `json:"component" query:"component"`
	Range       string `json:"range" query:"range"`
	Area        string `json:"area" query:"area"`
	Duration    string `json:"duration" query:"duration"`
	Target      string `json:"target" query:"target"`
	Rank        uint8  `json:"rank" query:"rank"`
	Ritual      bool   `json:"ritual" query:"ritual"`
	School      School `json:"school" query:"school"`
	Cast        string `json:"cast" query:"cast"`
	TraitsID    []uint `json:"traits_id" query:"traits_id"`
	TraditionID []uint `json:"tradition_id" query:"tradition_id"`
}

type SpellExternal struct {
	ID          uint     `json:"id" query:"id"`
	Name        string   `json:"name" query:"name"`
	Description string   `json:"description" query:"description"`
	Component   string   `json:"component" query:"component"`
	Range       string   `json:"range" query:"range"`
	Area        string   `json:"area" query:"area"`
	Duration    string   `json:"duration" query:"duration"`
	Target      string   `json:"target" query:"target"`
	Rank        uint8    `json:"rank" query:"rank"`
	Ritual      bool     `json:"ritual" query:"ritual"`
	School      School   `json:"school" query:"school"`
	Cast        string   `json:"cast" query:"cast"`
	Traits      []string `json:"traits" query:"traits"`
	Tradition   []string `json:"tradition" query:"tradition"`
}
