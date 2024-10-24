package api

import (
	"kingdom/model"
)

func (a *CharacterApi) CreateCharacterInfo(characterID uint, strength uint8) {
	characterInfo := &model.CharacterInfo{
		CharacterID: characterID,
		MaxBulk:     float64(10 + uint8(strength/2)),
	}
	err := a.DB.CreateCharacterInfo(characterInfo)
	if err != nil {
		return
	}
}

func (a *CharacterItemApi) UpdateCharacterBulk(characterId uint, bulk float64, isRemove bool) {
	characterInfo, _ := a.DB.GetCharacterInfoByID(characterId)
	if isRemove {
		characterInfo.Bulk -= bulk
	} else {
		characterInfo.Bulk += bulk
	}
	err := a.DB.UpdateCharacterInfo(characterInfo)
	if err != nil {
		return
	}
}
