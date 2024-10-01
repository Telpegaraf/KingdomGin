package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type GodDatabase interface {
	GetGodByID(id uint) (*model.God, error)
	CreateGod(god *model.God) error
}

type GodApi struct {
	DB GodDatabase
}

// GetGodById godoc
//
// @Summary returns God by id
// @Description Retrieve God details using its ID
// @Tags God
// @Accept json
// @Produce json
// @Param id path int true "god id"
// @Success 200 {object} model.God "God details"
// @Failure 404 {string} string "God not found"
// @Router /god/{id} [get]
func (a *GodApi) GetGodById(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		god, err := a.DB.GetGodByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if god != nil {
			ctx.JSON(http.StatusOK, god)
		} else {
			ctx.JSON(404, gin.H{"error": "God not found"})
		}
	})
}
