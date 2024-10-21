package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterBoostDatabase interface {
	GetCharacterBoostByID(id uint) (*model.CharacterBoost, error)
	UpdateCharacterBoost(CharacterBoosts *model.CharacterBoost) error
}

type CharacterBoostApi struct {
	DB CharacterBoostDatabase
}

func (a *CharacterApi) CreateCharacterBoost(ctx *gin.Context,
	characterID uint,
	race *model.Race) {

	internal := &model.CharacterBoost{
		CharacterID:   characterID,
		AncestryBoost: race.AbilityBoost,
	}
	if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacterBoost(internal)); !success {
		return
	}
}

// GetCharacterBoostByID godoc
//
// @Summary Returns CharacterBoost by id
// @Description Permissions for auth user or admin
// @Tags CharacterBoost
// @Accept json
// @Produce json
// @Param id path int true "character_id"
// @Success 200 {object} model.CharacterBoostExternal "CharacterBoost details"
// @Failure 404 {string} string "CharacterBoost not found"
// @Router /character_boost/{id} [get]
func (a *CharacterBoostApi) GetCharacterBoostByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		CharacterBoost, err := a.DB.GetCharacterBoostByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if CharacterBoost != nil {
			ctx.JSON(http.StatusOK, ToExternalCharacterBoost(CharacterBoost))
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "CharacterBoost doesn't exist"})
		}
	})
}

// UpdateCharacterBoost Updates CharacterBoost by ID
//
// @Summary Updates CharacterBoost by ID or nil
// @Description Permissions for Character's User or Admin
// @Tags CharacterBoost
// @Accept json
// @Produce json
// @Param id path int true "CharacterBoost id"
// @Param CharacterBoost body model.UpdateCharacterBoost true "CharacterBoost data"
// @Success 200 {object} model.CharacterBoostExternal "CharacterBoost details"
// @Failure 404 {string} string "CharacterBoost doesn't exist"
// @Router /character_boost/{id} [patch]
func (a *CharacterBoostApi) UpdateCharacterBoost(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var CharacterBoost *model.UpdateCharacterBoost
		if err := ctx.ShouldBindJSON(&CharacterBoost); err == nil {
			oldCharacterBoost, err := a.DB.GetCharacterBoostByID(id)
			if success := SuccessOrAbort(ctx, 404, err); !success {
				return
			}
			if oldCharacterBoost != nil {
				internal := &model.CharacterBoost{
					ID:          oldCharacterBoost.ID,
					CharacterID: id,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateCharacterBoost(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalCharacterBoost(internal))
			}
		}
	})
}

func ToExternalCharacterBoost(CharacterBoost *model.CharacterBoost) *model.CharacterBoostExternal {
	return &model.CharacterBoostExternal{
		ID:          CharacterBoost.ID,
		CharacterID: CharacterBoost.CharacterID,
	}
}
