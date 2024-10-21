package api

import "kingdom/model"

func (a *CharacterApi) CreateCharacterFeat(characterID uint, background *model.Background) {
	characterFeat := &model.CharacterFeat{
		CharacterID: characterID,
		FeatID:      background.FeatID,
	}
	a.DB.CreateCharacterFeat(characterFeat)
}
