package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

type AncestryDatabase interface {
	CreateAncestry(ancestry *model.Ancestry) error
}

type AncestryApi struct {
	DB AncestryDatabase
}

func (a *AncestryApi) CreateAncestry(ctx *gin.Context) {
	ancestry := model.Ancestry{}
}
