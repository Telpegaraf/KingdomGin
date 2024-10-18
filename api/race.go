package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type RaceDatabase interface {
	CreateRace(race *model.Race) error
	GetRaceByID(id uint) (*model.Race, error)
	GetRaces() ([]*model.Race, error)
	DeleteRaceByID(id uint) error
	UpdateRace(race *model.Race) error
}

type RaceApi struct {
	DB RaceDatabase
}

// CreateRace godoc
//
// @Summary Create and returns Race or nil
// @Description Permissions for Admin
// @Tags Race
// @Accept json
// @Produce json
// @Param race body model.RaceCreate true "Feat data"
// @Success 201 {object} model.RaceExternal "Feat details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /race [post]
func (a *RaceApi) CreateRace(ctx *gin.Context) {
	race := &model.RaceCreate{}
	if err := ctx.Bind(race); err == nil {
		internal := &model.Race{
			Name:         race.Name,
			Description:  race.Description,
			HitPoint:     race.HitPoint,
			Speed:        race.Speed,
			Size:         race.Size,
			AbilityBoost: race.AbilityBoost,
			Language:     race.Language,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateRace(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalRace(internal))
	}
}

// GetRaces godoc
//
// @Summary Returns all Races
// @Description Return all Races
// @Tags Race
// @Accept json
// @Produce json
// @Success 200 {object} model.Race "Race details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /race [get]
func (a *RaceApi) GetRaces(ctx *gin.Context) {
	races, err := a.DB.GetRaces()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.RaceExternal
	for _, race := range races {
		externalRace := ToExternalRace(race)
		resp = append(resp, externalRace)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetRaceByID godoc
//
// @Summary Returns Race by ID
// @Description Permissions for auth users
// @Tags Race
// @Accept json
// @Produce json
// @Param id path int true "Race id"
// @Success 200 {object} model.Race "Race details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /race/{id} [get]
func (a *RaceApi) GetRaceByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		race, err := a.DB.GetRaceByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalRace(race))
	})
}

// UpdateRace Updates Race by ID
//
// @Summary Updates Race by ID or nil
// @Description Permissions for Admin
// @Tags Race
// @Accept json
// @Produce json
// @Param id path int true "Race id"
// @Param race body model.RaceUpdate true "Race data"
// @Success 200 {object} model.Race "Race details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Race doesn't exist"
// @Router /race/{id} [patch]
func (a *RaceApi) UpdateRace(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var race model.RaceUpdate
		if err := ctx.ShouldBindJSON(&race); err == nil {
			oldRace, err := a.DB.GetRaceByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			internalRace := &model.Race{
				ID:           oldRace.ID,
				Name:         race.Name,
				Description:  race.Description,
				HitPoint:     race.HitPoint,
				Speed:        race.Speed,
				Size:         race.Size,
				AbilityBoost: race.AbilityBoost,
				Language:     race.Language,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateRace(internalRace)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			newRace, _ := a.DB.GetRaceByID(id)
			ctx.JSON(http.StatusOK, ToExternalRace(newRace))
		}
	})
}

// DeleteRace Deletes Race by ID
//
// @Summary Deletes Race by ID or returns nil
// @Description Permissions for Admin
// @Tags Race
// @Accept json
// @Produce json
// @Param id path int true "Race id"
// @Success 204
// @Failure 404 {string} string "Domain doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /Race/{id} [delete]
func (a *RaceApi) DeleteRace(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		race, err := a.DB.GetRaceByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if race != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteRaceByID(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalRace(race *model.Race) *model.RaceExternal {
	return &model.RaceExternal{
		ID:           race.ID,
		Name:         race.Name,
		Description:  race.Description,
		HitPoint:     race.HitPoint,
		Speed:        race.Speed,
		Size:         race.Size,
		AbilityBoost: race.AbilityBoost,
		Language:     race.Language,
	}
}
