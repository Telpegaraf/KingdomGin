package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type FeatDatabase interface {
	CreateFeat(feat *model.Feat) error
}

type FeatAPI struct {
	DB FeatDatabase
}

func (a *FeatAPI) CreateFeat(ctx *gin.Context) {
	feat := &model.Feat{}
	if err := ctx.Bind(&feat); err == nil {
		internal := &model.Feat{
			Name:        feat.Name,
			Description: feat.Description,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateFeat(internal)); !success {
			return
		}
	}
	ctx.JSON(http.StatusCreated, feat)
}
