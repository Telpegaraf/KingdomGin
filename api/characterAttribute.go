package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type AttributeDatabase interface {
	GetAttributeByID(id uint) (*model.Attribute, error)
	UpdateAttribute(attributes *model.Attribute) error
}

type AttributeApi struct {
	DB AttributeDatabase
}

func (a *CharacterApi) CreateAttribute(
	ctx *gin.Context,
	characterID uint,
	race *model.Race,
) {
	internal := &model.Attribute{
		CharacterID: characterID,
	}
	switch race.AttributeFlaw {
	case model.Strength:
		internal.Strength = 8
	case model.Dexterity:
		internal.Dexterity = 8
	case model.Constitution:
		internal.Constitution = 8
	case model.Intelligence:
		internal.Intelligence = 8
	case model.Wisdom:
		internal.Wisdom = 8
	case model.Charisma:
		internal.Charisma = 8
	}
	a.DB.CreateAttribute(internal)
}

// GetAttributeByID godoc
//
// @Summary Returns attribute by id
// @Description Permissions for auth user or admin
// @Tags Attribute
// @Accept json
// @Produce json
// @Param id path int true "character_id"
// @Success 200 {object} model.AttributeExternal "Attribute details"
// @Failure 404 {string} string "Attribute not found"
// @Router /attribute/{id} [get]
func (a *AttributeApi) GetAttributeByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		attribute, err := a.DB.GetAttributeByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if attribute != nil {
			ctx.JSON(http.StatusOK, ToExternalAttribute(attribute))
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Attribute doesn't exist"})
		}
	})
}

// UpdateAttribute Updates Attribute by ID
//
// @Summary Updates Attribute by ID or nil
// @Description Permissions for Character's User or Admin
// @Tags Attribute
// @Accept json
// @Produce json
// @Param id path int true "Attribute id"
// @Param attribute body model.UpdateAttribute true "Attribute data"
// @Success 200 {object} model.AttributeExternal "Attribute details"
// @Failure 404 {string} string "Attribute doesn't exist"
// @Router /attribute/{id} [patch]
func (a *AttributeApi) UpdateAttribute(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var attribute *model.UpdateAttribute
		if err := ctx.ShouldBindJSON(&attribute); err == nil {
			oldAttribute, err := a.DB.GetAttributeByID(id)
			if success := SuccessOrAbort(ctx, 404, err); !success {
				return
			}
			if oldAttribute != nil {
				internal := &model.Attribute{
					ID:           oldAttribute.ID,
					Strength:     *attribute.Strength,
					Dexterity:    *attribute.Dexterity,
					Constitution: *attribute.Constitution,
					Intelligence: *attribute.Intelligence,
					Wisdom:       *attribute.Wisdom,
					Charisma:     *attribute.Charisma,
					CharacterID:  id,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateAttribute(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalAttribute(internal))
			}
		}
	})
}

func ToExternalAttribute(attribute *model.Attribute) *model.AttributeExternal {
	return &model.AttributeExternal{
		ID:           attribute.ID,
		CharacterID:  attribute.CharacterID,
		Strength:     attribute.Strength,
		Dexterity:    attribute.Dexterity,
		Constitution: attribute.Constitution,
		Intelligence: attribute.Intelligence,
		Wisdom:       attribute.Wisdom,
		Charisma:     attribute.Charisma,
	}
}
