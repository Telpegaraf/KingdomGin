package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type RaceDatabase interface {
	CreateRace(race *model.Race) error
}

type RaceApi struct {
	DB RaceDatabase
}

func (a *RaceApi) CreateRace(ctx *gin.Context) {
	race := &model.Race{}
	if err := ctx.Bind(race); err == nil {
		internal := &model.Race{
			Name:         race.Name,
			Description:  race.Description,
			HitPoints:    race.HitPoints,
			Speed:        race.Speed,
			Size:         race.Size,
			AbilityBoost: race.AbilityBoost,
			Language:     race.Language,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateRace(internal)); !success {
			return
		}
	}
	ctx.JSON(http.StatusCreated, race)
}
