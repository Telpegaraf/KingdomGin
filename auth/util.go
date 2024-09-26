package auth

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

func TryGetUserID(ctx *gin.Context) *uint {
	user := ctx.MustGet("user").(*model.User)
	if user == nil {
		userID := ctx.MustGet("userid").(uint)
		if userID == 0 {
			return nil
		}
		return &userID
	}

	return &user.ID
}
