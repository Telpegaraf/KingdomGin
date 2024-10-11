package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

func (a *CharacterApi) CreateAttribute(ctx *gin.Context, characterID uint) {
	internal := &model.Attribute{
		CharacterID: characterID,
	}
	if success := SuccessOrAbort(ctx, 500, a.DB.CreateAttribute(internal)); !success {
		return
	}
}
