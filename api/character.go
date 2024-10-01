package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterDatabase interface {
	GetCharacterByID(id uint) (*model.Character, error)
	CreateCharacter(character *model.Character) error
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
			ctx.JSON(404, gin.H{"error": "Character not found"})
		}
	})
}

// CreateCharacter godoc
//
// @Summary Create and returns character or nil
// @Description Create new character
// @Tags Character
// @Accept json
// @Produce json
// @Param user body model.CreateCharacter true "Character data"
// @Success 201 {object} model.CharacterExternal "Character details"
// @Failure 401 {string} string "Unauthorized"
// @Router /character [post]
func (a *CharacterApi) CreateCharacter(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDFloat64, ok := userID.(float64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userID type"})
		return
	}
	userIDInt := uint(userIDFloat64)

	character := &model.Character{}
	if err := ctx.Bind(character); err == nil {
		internal := &model.Character{
			Name:     character.Name,
			Alias:    character.Alias,
			LastName: character.LastName,
			UserID:   userIDInt,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacter(internal)); !success {
			return
		}
	}
}

func ToExternalCharacter(internal *model.Character) *model.CharacterExternal {
	return &model.CharacterExternal{
		ID:       internal.ID,
		Name:     internal.Name,
		Alias:    internal.Alias,
		LastName: internal.LastName,
		UserID:   internal.UserID,
	}
}
