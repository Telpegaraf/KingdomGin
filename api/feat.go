package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
	"strconv"
)

type FeatDatabase interface {
	CreateFeat(feat *model.Feat) error
	GetFeatByID(id uint) (*model.Feat, error)
	GetFeats(limit int, offset int) (*[]model.Feat, error)
	DeleteFeat(id uint) error
	UpdateFeat(feat *model.Feat) error
}

type FeatAPI struct {
	DB FeatDatabase
}

// CreateFeat godoc
//
// @Summary Create and returns Feat or nil
// @Description Permissions for Admin
// @Tags Feat
// @Accept json
// @Produce json
// @Param feat body model.CreateFeat true "Feat data"
// @Success 201 {object} model.FeatExternal "Feat details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /feat [post]
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

// GetFeats godoc
//
// @Summary Returns all Feats
// @Description Return all Feats
// @Tags Feat
// @Accept json
// @Produce json
// @Param limit query int false "Limit for pagination"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} model.FeatExternal "Feat details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /feat [get]
func (a *FeatAPI) GetFeats(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "0")
	offset := ctx.DefaultQuery("offset", "0")
	limitInt, err := strconv.Atoi(limit)
	offsetInt, err := strconv.Atoi(offset)
	feats, err := a.DB.GetFeats(limitInt, limitInt*offsetInt)
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.FeatExternal
	for _, feat := range *feats {
		externalFeat := ToExternalFeat(&feat)
		resp = append(resp, externalFeat)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetFeatByID godoc
//
// @Summary Returns Feat by ID
// @Description Permissions for auth users
// @Tags Feat
// @Accept json
// @Produce json
// @Param id path int true "Feat id"
// @Success 200 {object} model.FeatExternal "Feat details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /feat/{id} [get]
func (a *FeatAPI) GetFeatByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		Feat, err := a.DB.GetFeatByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalFeat(Feat))
	})
}

// UpdateFeat Updates Feat by ID
//
// @Summary Updates Feat by ID or nil
// @Description Permissions for Admin
// @Tags Feat
// @Accept json
// @Produce json
// @Param id path int true "Feat id"
// @Param Feat body model.UpdateFeat true "Feat data"
// @Success 200 {object} model.FeatExternal "Feat details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Feat doesn't exist"
// @Router /feat/{id} [patch]
func (a *FeatAPI) UpdateFeat(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var Feat model.UpdateFeat
		if err := ctx.ShouldBindJSON(&Feat); err == nil {
			oldFeat, err := a.DB.GetFeatByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			internalFeat := &model.Feat{
				ID:          oldFeat.ID,
				Name:        Feat.Name,
				Description: Feat.Description,
				Level:       Feat.Level,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateFeat(internalFeat)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			newFeat, _ := a.DB.GetFeatByID(id)
			ctx.JSON(http.StatusOK, ToExternalFeat(newFeat))
		}
	})
}

// DeleteFeat Deletes Feat by ID
//
// @Summary Deletes Feat by ID or returns nil
// @Description Permissions for Admin
// @Tags Feat
// @Accept json
// @Produce json
// @Param id path int true "Feat id"
// @Success 204
// @Failure 404 {string} string "Domain doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /feat/{id} [delete]
func (a *FeatAPI) DeleteFeat(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		feat, err := a.DB.GetFeatByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if feat != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteFeat(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalFeat(Feat *model.Feat) *model.FeatExternal {
	return &model.FeatExternal{
		ID:                  Feat.ID,
		Name:                Feat.Name,
		Description:         Feat.Description,
		Level:               Feat.Level,
		PrerequisiteSkillID: Feat.PrerequisiteSkillID,
		PrerequisiteMastery: Feat.PrerequisiteMastery,
		PrerequisiteFeat:    Feat.PrerequisiteFeat,
		Traits:              Feat.Traits,
	}
}
