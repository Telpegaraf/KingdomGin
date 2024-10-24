package test

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

func WithUser(ctx *gin.Context, userID uint) {
	ctx.Set("user", &model.User{ID: userID})
	ctx.Set("userID", userID)
}
