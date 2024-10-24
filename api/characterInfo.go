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

func (a *CharacterItemApi) UpdateCharacterBulk(characterId uint, bulk float64) {
	characterInfo, _ := a.DB.GetCharacterInfoByID(characterId)
	characterInfo.Bulk += bulk
	err := a.DB.UpdateCharacterInfo(characterInfo)
	if err != nil {
		return
	}
}
