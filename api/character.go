package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterDatabase interface {
	GetCharacterByID(id uint) (*model.Character, error)
}

type CharacterApi struct {
	DB CharacterDatabase
}

// GetCharacterByID godoc
//
// @Summary returns Character by id
// @Description Retrieve Character details using its ID
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "character id"
// @Success 200 {object} model.CharacterExternal "character details"
// @Failure 404 {string} string "Character not found"
// @Router /character/{id} [get]
func (a *CharacterApi) GetCharacterByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		character, err := a.DB.GetCharacterByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if character != nil {
			ctx.JSON(http.StatusOK, ToExternalCharacter(character))
		} else {
			ctx.JSON(404, errors.New("Character Not Found"))
		}
	})
}

func ToExternalCharacter(internal *model.Character) *model.CharacterExternal {
	return &model.CharacterExternal{
		Name:     internal.Name,
		Alias:    internal.Alias,
		LastName: internal.LastName,
	}
}
