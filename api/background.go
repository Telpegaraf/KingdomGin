package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type BackgroundDatabase interface {
	GetBackgroundByID(id uint) (*model.Background, error)
	CreateBackground(Background *model.Background) error
	GetBackgrounds() ([]*model.Background, error)
	UpdateBackground(Background *model.Background) error
	DeleteBackground(id uint) error
}

type BackgroundApi struct {
	DB BackgroundDatabase
}

// CreateBackground godoc
//
// @Summary Create and returns Background or nil
// @Description Permissions for Admin
// @Tags Background
// @Accept json
// @Produce json
// @Param Background body model.BackgroundCreate true "Background data"
// @Success 201 {object} model.BackgroundExternal "Background details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /background [post]
func (a *BackgroundApi) CreateBackground(ctx *gin.Context) {
	background := &model.BackgroundCreate{}
	if err := ctx.ShouldBindJSON(background); err == nil {
		internal := &model.Background{
			Name:          background.Name,
			Description:   background.Description,
			FeatID:        background.FeatID,
			FirstSkillID:  background.FirstSkillID,
			SecondSkillID: background.SecondSkillID,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateBackground(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToBackgroundExternal(internal))
	}
}

// GetBackgroundByID godoc
//
// @Summary Returns Background by id
// @Description Retrieve Background details using its ID
// @Tags Background
// @Accept json
// @Produce json
// @Param id path int true "Background id"
// @Success 200 {object} model.Background "Background details"
// @Failure 404 {string} string "Background not found"
// @Router /background/{id} [get]
func (a *BackgroundApi) GetBackgroundByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		background, err := a.DB.GetBackgroundByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if background != nil {
			ctx.JSON(http.StatusOK, ToBackgroundExternal(background))
		}
	})
}

// GetBackgrounds godoc
//
// @Summary Returns all Backgrounds
// @Description Return all Backgrounds
// @Tags Background
// @Accept json
// @Produce json
// @Success 200 {object} model.Background "Background details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /background [get]
func (a *BackgroundApi) GetBackgrounds(ctx *gin.Context) {
	backgrounds, err := a.DB.GetBackgrounds()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.BackgroundExternal
	for _, background := range backgrounds {
		resp = append(resp, ToBackgroundExternal(background))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateBackground Updates Background by ID
//
// @Summary Updates Background by ID or nil
// @Description Permissions for Admin
// @Tags Background
// @Accept json
// @Produce json
// @Param id path int true "Background id"
// @Param background body model.BackgroundUpdate true "Background data"
// @Success 200 {object} model.BackgroundExternal "Background details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Background doesn't exist"
// @Router /background/{id} [patch]
func (a *BackgroundApi) UpdateBackground(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var background *model.BackgroundUpdate
		if err := ctx.Bind(&background); err == nil {
			oldBackground, err := a.DB.GetBackgroundByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldBackground != nil {
				internal := &model.Background{
					ID:            oldBackground.ID,
					Name:          background.Name,
					Description:   background.Description,
					FeatID:        background.FeatID,
					FirstSkillID:  background.FirstSkillID,
					SecondSkillID: background.SecondSkillID,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateBackground(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToBackgroundExternal(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Background doesn't exist"})
		}
	})
}

// DeleteBackground Deletes Background by ID
//
// @Summary Deletes Background by ID or returns nil
// @Description Permissions for Admin
// @Tags Background
// @Accept json
// @Produce json
// @Param id path int true "Background id"
// @Success 204
// @Failure 404 {string} string "Background doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /background/{id} [delete]
func (a *BackgroundApi) DeleteBackground(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		background, err := a.DB.GetBackgroundByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if background != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteBackground(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Background was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Background doesn't exist"})
		}
	})
}

func ToBackgroundExternal(Background *model.Background) *model.BackgroundExternal {
	return &model.BackgroundExternal{
		ID:            Background.ID,
		Name:          Background.Name,
		Description:   Background.Description,
		FeatID:        Background.FeatID,
		FirstSkillID:  Background.FirstSkillID,
		SecondSkillID: Background.SecondSkillID,
	}
}
