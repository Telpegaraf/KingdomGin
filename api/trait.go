package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type TraitDatabase interface {
	GetTraitByID(id uint) (*model.Trait, error)
	CreateTrait(Trait *model.Trait) error
	GetTraits() ([]*model.Trait, error)
	UpdateTrait(Trait *model.Trait) error
	DeleteTrait(id uint) error
}

type TraitApi struct {
	DB TraitDatabase
}

// CreateTrait godoc
//
// @Summary Create and returns Trait or nil
// @Description Permissions for Admin
// @Tags Trait
// @Accept json
// @Produce json
// @Param trait body model.CreateTrait true "Trait data"
// @Success 201 {object} model.TraitExternal "Trait details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /trait [post]
func (a *TraitApi) CreateTrait(ctx *gin.Context) {
	trait := &model.CreateTrait{}
	if err := ctx.ShouldBindJSON(trait); err == nil {
		internal := &model.Trait{
			Name:        trait.Name,
			Description: trait.Description,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateTrait(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToTraitExternal(internal))
	}
}

// GetTraitByID godoc
//
// @Summary Returns Trait by id
// @Description Retrieve Trait details using its ID
// @Tags Trait
// @Accept json
// @Produce json
// @Param id path int true "Trait id"
// @Success 200 {object} model.TraitExternal "Trait details"
// @Failure 404 {string} string "Trait not found"
// @Router /trait/{id} [get]
func (a *TraitApi) GetTraitByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		trait, err := a.DB.GetTraitByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if trait != nil {
			ctx.JSON(http.StatusOK, ToTraitExternal(trait))
		}
	})
}

// GetTraits godoc
//
// @Summary Returns all Traits
// @Description Return all Traits
// @Tags Trait
// @Accept json
// @Produce json
// @Success 200 {object} model.TraitExternal "Trait details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /trait [get]
func (a *TraitApi) GetTraits(ctx *gin.Context) {
	traits, err := a.DB.GetTraits()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.TraitExternal
	for _, trait := range traits {
		resp = append(resp, ToTraitExternal(trait))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateTrait Updates Trait by ID
//
// @Summary Updates Trait by ID or nil
// @Description Permissions for Admin
// @Tags Trait
// @Accept json
// @Produce json
// @Param id path int true "Trait id"
// @Param character body model.UpdateTrait true "Trait data"
// @Success 200 {object} model.TraitExternal "Trait details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Trait doesn't exist"
// @Router /trait/{id} [patch]
func (a *TraitApi) UpdateTrait(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var trait *model.UpdateTrait
		if err := ctx.Bind(&trait); err == nil {
			oldTrait, err := a.DB.GetTraitByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldTrait != nil {
				internal := &model.Trait{
					ID:          oldTrait.ID,
					Name:        trait.Name,
					Description: trait.Description,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateTrait(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToTraitExternal(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Trait doesn't exist"})
		}
	})
}

// DeleteTrait Deletes Trait by ID
//
// @Summary Deletes Trait by ID or returns nil
// @Description Permissions for Admin
// @Tags Trait
// @Accept json
// @Produce json
// @Param id path int true "Trait id"
// @Success 204
// @Failure 404 {string} string "Trait doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /trait/{id} [delete]
func (a *TraitApi) DeleteTrait(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		trait, err := a.DB.GetTraitByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if trait != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteTrait(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Trait was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Trait doesn't exist"})
		}
	})
}

func ToTraitExternal(trait *model.Trait) *model.TraitExternal {
	return &model.TraitExternal{
		ID:          trait.ID,
		Name:        trait.Name,
		Description: trait.Description,
	}
}
