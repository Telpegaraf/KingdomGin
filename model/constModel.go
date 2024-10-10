package model

type Health int8
type MasteryLevel string
type Size string

const (
	Six    Health = 6
	Eight  Health = 8
	Ten    Health = 10
	Twelve Health = 12
)

const (
	None   MasteryLevel = "None"
	Train  MasteryLevel = "Train"
	Expert MasteryLevel = "Expert"
	Master MasteryLevel = "Master"
	Legend MasteryLevel = "Legend"
)

const (
	Tiny       Size = "Tiny"
	Small      Size = "Small"
	Medium     Size = "Medium"
	Large      Size = "Large"
	Huge       Size = "Huge"
	Gargantuan Size = "Gargantuan"
)
