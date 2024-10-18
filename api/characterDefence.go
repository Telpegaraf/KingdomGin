package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

type CharacterDefenceDatabase interface {
	GetCharacterBoostByID(id uint) (*model.CharacterBoost, error)
	UpdateCharacterBoost(CharacterBoosts *model.CharacterBoost) error
}

type CharacterDefenceApi struct {
	DB CharacterBoostDatabase
}

func (a *CharacterApi) CreateCharacterDefence(ctx *gin.Context, character *model.Character) {
	race, err := a.DB.GetRaceByID(character.RaceID)
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	characterClass, err := a.DB.GetCharacterClassByID(character.BackgroundID)
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	internal := &model.CharacterDefence{
		CharacterID: character.ID,
		HitPoint:    race.HitPoint + characterClass.HitPoint,
	}
	if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacterDefence(internal)); !success {
		return
	}
}
