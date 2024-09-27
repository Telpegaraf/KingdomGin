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
	user := ctx.MustGet("user").(*model.User)
	if user == nil {
		userID := ctx.MustGet("userID").(uint)
		if userID == 0 {
			return nil
		}
		return &userID
	}

	return &user.ID
}

func GetTokenID(ctx *gin.Context) string {
	return ctx.MustGet("tokenID").(string)
}
