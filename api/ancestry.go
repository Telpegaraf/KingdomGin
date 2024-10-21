package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type AncestryDatabase interface {
	CreateAncestry(ancestry *model.Ancestry) error
	GetAncestryByID(id uint) (*model.Ancestry, error)
	GetAncestries() ([]*model.Ancestry, error)
	UpdateAncestry(ancestry *model.Ancestry) error
	DeleteAncestry(id uint) error
}

type AncestryApi struct {
	DB AncestryDatabase
}

// CreateAncestry godoc
//
// @Summary Create and returns Ancestry or nil
// @Description Permissions for Admin
// @Tags Ancestry
// @Accept json
// @Produce json
// @Param ancestry body model.AncestryCreate true "Feat data"
// @Success 201 {object} model.AncestryExternal "Feat details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /ancestry [post]
func (a *AncestryApi) CreateAncestry(ctx *gin.Context) {
	ancestry := &model.AncestryCreate{}
	if err := ctx.Bind(&ancestry); err == nil {
		internal := &model.Ancestry{
			Name:        ancestry.Name,
			Description: ancestry.Description,
			RaceID:      ancestry.RaceID,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateAncestry(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalAncestry(internal))
	}
}

// GetAncestries godoc
//
// @Summary Returns all Ancestries
// @Description Return all GetAncestries
// @Tags Ancestry
// @Accept json
// @Produce json
// @Success 200 {object} model.Ancestry "Ancestry details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /ancestry [get]
func (a *AncestryApi) GetAncestries(ctx *gin.Context) {
	ancestries, err := a.DB.GetAncestries()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.AncestryExternal
	for _, ancestry := range ancestries {
		externalAncestry := ToExternalAncestry(ancestry)
		resp = append(resp, externalAncestry)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetAncestryByID godoc
//
// @Summary Returns Ancestry by ID
// @Description Permissions for auth users
// @Tags Ancestry
// @Accept json
// @Produce json
// @Param id path int true "Ancestry id"
// @Success 200 {object} model.Ancestry "Ancestry details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /ancestry/{id} [get]
func (a *AncestryApi) GetAncestryByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		ancestry, err := a.DB.GetAncestryByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalAncestry(ancestry))
	})
}

// UpdateAncestry Updates Ancestry by ID
//
// @Summary Updates Ancestry by ID or nil
// @Description Permissions for Admin
// @Tags Ancestry
// @Accept json
// @Produce json
// @Param id path int true "Ancestry id"
// @Param Ancestry body model.AncestryUpdate true "Ancestry data"
// @Success 200 {object} model.Ancestry "Ancestry details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Ancestry doesn't exist"
// @Router /ancestry/{id} [patch]
func (a *AncestryApi) UpdateAncestry(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var ancestry model.AncestryUpdate
		if err := ctx.ShouldBindJSON(&ancestry); err == nil {
			oldAncestry, err := a.DB.GetAncestryByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			internalAncestry := &model.Ancestry{
				ID:          oldAncestry.ID,
				Name:        ancestry.Name,
				Description: ancestry.Description,
				RaceID:      ancestry.RaceID,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateAncestry(internalAncestry)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			newAncestry, _ := a.DB.GetAncestryByID(id)
			ctx.JSON(http.StatusOK, ToExternalAncestry(newAncestry))
		}
	})
}

// DeleteAncestry Deletes Ancestry by ID
//
// @Summary Deletes Ancestry by ID or returns nil
// @Description Permissions for Admin
// @Tags Ancestry
// @Accept json
// @Produce json
// @Param id path int true "Ancestry id"
// @Success 204
// @Failure 404 {string} string "Domain doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /ancestry/{id} [delete]
func (a *AncestryApi) DeleteAncestry(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		ancestry, err := a.DB.GetAncestryByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if ancestry != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteAncestry(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalAncestry(Ancestry *model.Ancestry) *model.AncestryExternal {
	return &model.AncestryExternal{
		ID:          Ancestry.ID,
		Name:        Ancestry.Name,
		Description: Ancestry.Description,
		RaceID:      Ancestry.RaceID,
	}
}
