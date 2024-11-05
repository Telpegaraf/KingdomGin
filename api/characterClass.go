package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterClassDatabase interface {
	CreateCharacterClass(characterClass *model.CharacterClass) error
	GetCharacterClassByID(id uint) (*model.CharacterClass, error)
	GetCharacterClasses() (*[]model.CharacterClass, error)
	DeleteCharacterClass(id uint) error
	UpdateCharacterClass(class *model.CharacterClass) error
	GetTraditionByName(name string) (*model.Tradition, error)
}

type CharacterClassApi struct {
	DB CharacterClassDatabase
}

// CreateCharacterClass godoc
//
// @Summary Create and returns new Character Class or nil
// @Description Create new Character Class
// @Tags Character Class
// @Accept json
// @Produce json
// @Param god body model.CharacterClassCreate true "Character Class data"
// @Success 201 {object} model.CharacterClass "Character Class details"
// @Failure 401 {string} string "Unauthorized"
// @Router /class [post]
func (a *CharacterClassApi) CreateCharacterClass(ctx *gin.Context) {
	characterClass := &model.CharacterClassCreate{}
	if err := ctx.Bind(characterClass); err == nil {
		internal := &model.CharacterClass{
			Name:     characterClass.Name,
			HitPoint: characterClass.HitPoint,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacterClass(internal)); !success {
			return
		}
	}
}

// GetCharacterClassByID godoc
//
// @Summary Returns Character by id
// @Description Retrieve Character class details using its ID
// @Tags Character Class
// @Accept json
// @Produce json
// @Param id path int true "character id"
// @Success 200 {object} model.CharacterExternal "character details"
// @Failure 404 {string} string "Character not found"
// @Router /class/{id} [get]
func (a *CharacterClassApi) GetCharacterClassByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		character, err := a.DB.GetCharacterClassByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if character != nil {
			ctx.JSON(http.StatusOK, ToExternalCharacterClass(character))
		} else {
			ctx.JSON(404, gin.H{"error": "Character not found"})
		}
	})
}

// GetCharacterClasses godoc
//
// @Summary Returns all characters
// @Description Return all characters for current user
// @Tags Character Class
// @Accept json
// @Produce json
// @Success 200 {object} model.CharacterExternal "Character details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /class [get]
func (a *CharacterClassApi) GetCharacterClasses(ctx *gin.Context) {
	characters, err := a.DB.GetCharacterClasses()

	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.CharacterClassExternal
	for _, character := range *characters {
		resp = append(resp, ToExternalCharacterClass(&character))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateCharacterClass Updates Character by ID
//
// @Summary Updates Character by ID or nil
// @Description Permissions for Character's User or Admin
// @Tags Character Class
// @Accept json
// @Produce json
// @Param id path int true "Character id"
// @Param character body model.CharacterUpdate true "Character data"
// @Success 200 {object} model.CharacterExternal "Character details"
// @Failure 404 {string} string "Character doesn't exist"
// @Router /class/{id} [patch]
func (a *CharacterClassApi) UpdateCharacterClass(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var character *model.CharacterClassUpdate
		if err := ctx.Bind(&character); err == nil {
			oldCharacter, err := a.DB.GetCharacterClassByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldCharacter != nil {
				internal := &model.CharacterClass{
					ID:            oldCharacter.ID,
					Name:          character.Name,
					HitPoint:      character.HitPoint,
					Perception:    character.Perception,
					Fortitude:     character.Fortitude,
					Reflex:        character.Reflex,
					Will:          character.Will,
					UnarmedArmor:  character.UnarmedArmor,
					LightArmor:    character.LightArmor,
					MediumArmor:   character.MediumArmor,
					HeavyArmor:    character.HeavyArmor,
					UnArmedWeapon: character.UnArmedWeapon,
					CommonWeapon:  character.CommonWeapon,
					MartialWeapon: character.MartialWeapon,
					TraditionID:   &character.TraditionID,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateCharacterClass(internal)); success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalCharacterClass(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

// DeleteCharacterClass Deletes Character by ID
//
// @Summary Deletes Character by ID or returns nil
// @Description Permissions for Character's User or Admin
// @Tags Character Class
// @Accept json
// @Produce json
// @Param id path int true "Character id"
// @Success 204
// @Failure 404 {string} string "Character doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /class/{id} [delete]
func (a *CharacterClassApi) DeleteCharacterClass(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		character, err := a.DB.GetCharacterClassByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if character != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteCharacterClass(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalCharacterClass(character *model.CharacterClass) *model.CharacterClassExternal {
	return &model.CharacterClassExternal{
		ID:            character.ID,
		Name:          character.Name,
		HitPoint:      character.HitPoint,
		Perception:    character.Perception,
		Fortitude:     character.Fortitude,
		Reflex:        character.Reflex,
		Will:          character.Will,
		UnarmedArmor:  character.UnarmedArmor,
		LightArmor:    character.LightArmor,
		MediumArmor:   character.MediumArmor,
		HeavyArmor:    character.HeavyArmor,
		UnArmedWeapon: character.UnArmedWeapon,
		CommonWeapon:  character.CommonWeapon,
		MartialWeapon: character.MartialWeapon,
		TraditionID:   *character.TraditionID,
	}
}
