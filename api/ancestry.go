package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type AncestryDatabase interface {
	CreateAncestry(ancestry *model.Ancestry) error
}

type AncestryApi struct {
	DB AncestryDatabase
}

func (a *AncestryApi) CreateAncestry(ctx *gin.Context) {
	ancestry := &model.Ancestry{}
	if err := ctx.Bind(&ancestry); err == nil {
		internal := &model.Ancestry{
			Name:        ancestry.Name,
			Description: ancestry.Description,
			RaceID:      ancestry.RaceID,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateAncestry(internal)); !success {
			return
		}
	}
	ctx.JSON(http.StatusCreated, ancestry)
}
