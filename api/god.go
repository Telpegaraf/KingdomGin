package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type GodDatabase interface {
	GetGodByID(id uint) (*model.God, error)
	CreateGod(god *model.God) error
	GetGods() ([]*model.God, error)
	UpdateGod(god *model.God) error
	DeleteGod(id uint) error
}

type GodApi struct {
	DB GodDatabase
}

// CreateGod godoc
//
// @Summary Create and returns new God or nil
// @Description Create new God
// @Tags God
// @Accept json
// @Produce json
// @Param god body model.GodCreate true "God data"
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

// GetGodById godoc
//
// @Summary Returns God by id
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

// GetGods godoc
//
// @Summary Returns all gods
// @Description Return all gods and their domains
// @Tags God
// @Accept json
// @Produce json
// @Success 200 {object} model.God "Character details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /god [get]
func (a *GodApi) GetGods(ctx *gin.Context) {
	gods, err := a.DB.GetGods()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.God
	for _, god := range gods {
		resp = append(resp, god)
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateGod Updates God by ID
//
// @Summary Updates God by ID or nil
// @Description Updates God
// @Tags God
// @Accept json
// @Produce json
// @Param id path int true "God id"
// @Param character body model.GodUpdate true "God data"
// @Success 200 {object} model.God "God details"
// @Failure 404 {string} string "God doesn't exist"
// @Router /god/{id} [patch]
func (a *GodApi) UpdateGod(ctx *gin.Context) {
	var god *model.GodUpdate
	withID(ctx, "id", func(id uint) {
		oldGod, err := a.DB.GetGodByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if oldGod != nil {
			internal := &model.God{
				ID:              oldGod.ID,
				Name:            god.Name,
				Alias:           god.Alias,
				Edict:           god.Edict,
				Anathema:        god.Anathema,
				AreasOfInterest: god.AreasOfInterest,
				Worships:        god.Worships,
				SacredAnimals:   god.SacredAnimals,
				SacredColors:    god.SacredColors,
				ChosenWeapon:    god.ChosenWeapon,
				Alignment:       god.Alignment,
				Description:     god.Description,
				Domains:         god.Domains,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateGod(internal)); !success {
				return
			}
			ctx.JSON(http.StatusOK, internal)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "God doesn't exist"})
		}
	})
}

// DeleteGod Deletes God by ID
//
// @Summary Deletes God by ID or returns nil
// @Description Permissions for Auth user
// @Tags God
// @Accept json
// @Produce json
// @Param id path int true "God id"
// @Success 204
// @Failure 404 {string} string "God doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /god/{id} [delete]
func (a *GodApi) DeleteGod(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		god, err := a.DB.GetGodByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if god != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteGod(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"message": "God was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "God doesn't exist"})
		}
	})
}
