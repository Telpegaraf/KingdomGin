package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

type SlotDatabase interface {
}

type SlotApi struct {
	DB SlotDatabase
}

func (a *CharacterApi) CreateSlot(ctx *gin.Context, characterID uint) {
	internal := &model.Slot{
		CharacterID: characterID,
	}
	if success := SuccessOrAbort(ctx, 500, a.DB.CreateSlot(internal)); !success {
		return
	}
}
