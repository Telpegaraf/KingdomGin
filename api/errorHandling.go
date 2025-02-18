package api

import (
	"github.com/gin-gonic/gin"
)

func SuccessOrAbort(ctx *gin.Context, code int, err error) (success bool) {
	if err != nil {
		ctx.AbortWithError(code, err)
	}
	return err == nil
}
