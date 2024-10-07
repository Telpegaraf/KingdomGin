package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

func (a *CharacterApi) CreateStat(ctx *gin.Context, characterID uint) {
	internal := &model.Stat{
		CharacterID: characterID,
	}
	if success := SuccessOrAbort(ctx, 500, a.DB.CreateStat(internal)); !success {
		return
	}
}
