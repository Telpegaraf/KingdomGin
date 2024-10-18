package model

type Tradition string
type School string
type HitPoint string
type SquareSize string
type MasteryLevel string

const (
	Arcane Tradition = "Arcane"
	Divine Tradition = "Divine"
	Occult Tradition = "Occult"
	Primal Tradition = "Primal"
)

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
	Train  MasteryLevel = "Train"
	Expert MasteryLevel = "Expert"
	Master MasteryLevel = "Master"
	Legend MasteryLevel = "Legend"
)
