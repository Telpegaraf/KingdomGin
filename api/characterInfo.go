package api

import (
	"kingdom/model"
)

func (a *CharacterApi) CreateCharacterInfo(strength uint8) {
	characterInfo := &model.CharacterInfo{
		MaxBulk: float64(10 + uint8(strength/2)),
	}
	err := a.DB.CreateCharacterInfo(characterInfo)
	if err != nil {
		return
	}
}
