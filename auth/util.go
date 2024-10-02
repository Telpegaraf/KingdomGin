package auth

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

func RegisterAuthentication(ctx *gin.Context, user *model.User, userID uint, tokenID string) {
	ctx.Set("user", user)
	ctx.Set("userID", userID)
	ctx.Set("tokenID", tokenID)
}

func GetUserID(ctx *gin.Context) uint {
	id := TryGetUserID(ctx)
	if id == nil {
		panic("token and user may not be null")
	}
	return *id
}

func TryGetUserID(ctx *gin.Context) *uint {
	if userID, exists := ctx.Get("userID"); exists {
		if id, ok := userID.(uint); ok {
			return &id
		}
	}
	return nil
}

func GetTokenID(ctx *gin.Context) string {
	return ctx.MustGet("tokenID").(string)
}
