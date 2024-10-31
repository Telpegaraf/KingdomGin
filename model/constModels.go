package model

type School string
type HitPoint string
type SquareSize string
type MasteryLevel string
type Ability string
type Rarity string

const (
	Abjuration    School = "Abjuration"
	Conjuration   School = "Conjuration"
	Divination    School = "Divination"
	Enchantment   School = "Enchantment"
	Evocation     School = "Evocation"
	Illusion      School = "Illusion"
	Necromancy    School = "Necromancy"
	Transmutation School = "Transmutation"
)

const (
	Six    HitPoint = "Six"
	Eight  HitPoint = "Eight"
	Ten    HitPoint = "Ten"
	Twelve HitPoint = "Twelve"
)

const (
	Tiny       SquareSize = "Tiny"
	Small      SquareSize = "Small"
	Medium     SquareSize = "Medium"
	Large      SquareSize = "Large"
	Huge       SquareSize = "Huge"
	Gargantuan SquareSize = "Gargantuan"
)

const (
	None   MasteryLevel = "None"
	Train  MasteryLevel = "Trained"
	Expert MasteryLevel = "Expert"
	Master MasteryLevel = "Master"
	Legend MasteryLevel = "Legend"
)

const (
	Strength     Ability = "Strength"
	Dexterity    Ability = "Dexterity"
	Constitution Ability = "Constitution"
	Intelligence Ability = "Intelligence"
	Wisdom       Ability = "Wisdom"
	Charisma     Ability = "Charisma"
)

const (
	Common    Rarity = "Common"
	Uncommon  Rarity = "Uncommon"
	Rare      Rarity = "Rare"
	Legendary Rarity = "Legendary"
	Mythic    Rarity = "Mythic"
)
