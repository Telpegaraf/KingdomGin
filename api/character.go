package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/auth"
	"kingdom/model"
	"net/http"
)

type CharacterDatabase interface {
	GetCharacterByID(id uint) (*model.Character, error)
	CreateCharacter(character *model.Character) error
	GetCharacters(id uint) ([]*model.Character, error)
	UpdateCharacter(character *model.Character) error
	DeleteCharacterByID(id uint) error
	GetUserByID(id uint) (*model.User, error)
	CreateAttribute(stat *model.Attribute) error
}

type CharacterApi struct {
	DB CharacterDatabase
}

// GetCharacterByID godoc
//
// @Summary Returns Character by id
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
			ctx.JSON(http.StatusOK, character)
		} else {
			ctx.JSON(404, gin.H{"error": "Character not found"})
		}
	})
}

// GetCharacters godoc
//
// @Summary Returns all characters
// @Description Return all characters for current user
// @Tags Character
// @Accept json
// @Produce json
// @Success 200 {object} model.CharacterExternal "Character details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /character [get]
func (a *CharacterApi) GetCharacters(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	characters, err := a.DB.GetCharacters(userID.(uint))

	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.Character
	for _, character := range characters {
		resp = append(resp, character)
	}
	ctx.JSON(http.StatusOK, resp)
}

// CreateCharacter godoc
//
// @Summary Create and returns character or nil
// @Description Create new character
// @Tags Character
// @Accept json
// @Produce json
// @Param character body model.CreateCharacter true "Character data"
// @Success 201 {object} model.CharacterExternal "Character details"
// @Failure 401 {string} string "Unauthorized"
// @Router /character [post]
func (a *CharacterApi) CreateCharacter(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	character := &model.CreateCharacter{}
	if err := ctx.Bind(character); err == nil {
		internal := &model.Character{
			Name:     character.Name,
			Alias:    character.Alias,
			LastName: character.LastName,
			UserID:   userID.(uint),
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacter(internal)); !success {
			return
		}
		go func() {
			a.CreateAttribute(ctx, internal.ID)
		}()
	}
	ctx.JSON(http.StatusCreated, character)
}

// UpdateCharacter Updates Character by ID
//
// @Summary Updates Character by ID or nil
// @Description Permissions for Character's User or Admin
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "Character id"
// @Param character body model.CharacterUpdateExternal true "Character data"
// @Success 200 {object} model.CharacterExternal "Character details"
// @Failure 404 {string} string "Character doesn't exist"
// @Router /character/{id} [patch]
func (a *CharacterApi) UpdateCharacter(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))

	withID(ctx, "id", func(id uint) {
		var character *model.CharacterUpdateExternal
		if err := ctx.Bind(&character); err == nil {
			oldCharacter, err := a.DB.GetCharacterByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldCharacter != nil {
				if oldCharacter.UserID != user.ID && !user.Admin {
					ctx.JSON(http.StatusForbidden, gin.H{"error": "You can't access for this API"})
					return
				}
				internal := &model.Character{
					ID:       oldCharacter.ID,
					Name:     character.Name,
					Alias:    character.Alias,
					LastName: character.LastName,
					UserID:   oldCharacter.UserID,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateCharacter(internal)); success {
					return
				}
				ctx.JSON(http.StatusOK, internal)
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

// DeleteCharacter Deletes Character by ID
//
// @Summary Deletes Character by ID or returns nil
// @Description Permissions for Character's User or Admin
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "Character id"
// @Success 204
// @Failure 404 {string} string "Character doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /character/{id} [delete]
func (a *CharacterApi) DeleteCharacter(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))

	withID(ctx, "id", func(id uint) {
		character, err := a.DB.GetCharacterByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if character != nil {
			if character.UserID != user.ID && !user.Admin {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "You can't access for this API"})
				return
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteCharacterByID(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}
