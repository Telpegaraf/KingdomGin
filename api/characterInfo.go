package api

import "kingdom/model"

func (a *CharacterApi) CreateCharacterInfo(strength uint8) {
	characterInfo := &model.CharacterInfo{
		Bulk: float64(10 + uint8(strength/2)),
	}
	a.DB.CreateCharacterInfo(characterInfo)
}
