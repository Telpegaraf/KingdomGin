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

// CreateGod godoc
//
// @Summary Create and returns new God or nil
// @Description Create new God
// @Tags God
// @Accept json
// @Produce json
// @Param god body model.God true "God data"
// @Success 201 {object} model.God "God details"
// @Failure 401 {string} string "Unauthorized"
// @Router /god [post]
func (a *GodApi) CreateGod(ctx *gin.Context) {
	god := &model.God{}
	if err := ctx.Bind(god); err == nil {
		internal := &model.God{
			Name:            god.Name,
			Alias:           god.Alias,
			Edict:           god.Edict,
			Anathema:        god.Anathema,
			AreasOfInterest: god.AreasOfInterest,
			Temples:         god.Temples,
			Worships:        god.Worships,
			SacredAnimals:   god.SacredAnimals,
			SacredColors:    god.SacredColors,
			ChosenWeapon:    god.ChosenWeapon,
			Alignment:       god.Alignment,
			Description:     god.Description,
			Domains:         god.Domains,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateGod(internal)); !success {
			return
		}
	}
}
