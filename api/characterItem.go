package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterItemDatabase interface {
	CreateCharacterItem(characterItem *model.CharacterItem) error
	GetCharacterItemByID(id uint) (*model.CharacterItem, error)
	GetCharacterItems() ([]*model.CharacterItem, error)
	UpdateCharacterItem(item *model.CharacterItem) error
	DeleteCharacterItem(id uint) error
}

type CharacterItemApi struct {
	DB CharacterItemDatabase
}

// CreateCharacterItem godoc
//
// @Summary Create and returns CharacterItem or nil
// @Description Permissions for Admin
// @Tags CharacterItem
// @Accept json
// @Produce json
// @Param characterItem body model.CreateCharacterItem true "CharacterItem data"
// @Success 201 {object} model.CharacterItemExternal "CharacterItem details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /CharacterItem [post]
func (a *CharacterItemApi) CreateCharacterItem(ctx *gin.Context) {
	CharacterItem := &model.CreateCharacterItem{}
	if err := ctx.ShouldBindJSON(CharacterItem); err == nil {
		internal := &model.CharacterItem{
			CharacterId: CharacterItem.CharacterID,
			ItemId:      CharacterItem.ItemID,
			Quantity:    CharacterItem.Quantity,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacterItem(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalCharacterItem(internal))
	}
}

// GetCharacterItemByID godoc
//
// @Summary Returns CharacterItem by id
// @Description Retrieve CharacterItem details using its ID
// @Tags CharacterItem
// @Accept json
// @Produce json
// @Param id path int true "CharacterItem id"
// @Success 200 {object} model.CharacterItem "CharacterItem details"
// @Failure 404 {string} string "CharacterItem not found"
// @Router /CharacterItem/{id} [get]
func (a *CharacterItemApi) GetCharacterItemByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		CharacterItem, err := a.DB.GetCharacterItemByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if CharacterItem != nil {
			ctx.JSON(http.StatusOK, CharacterItem)
		}
	})
}

// GetCharacterItems godoc
//
// @Summary Returns all CharacterItems
// @Description Return all CharacterItems
// @Tags CharacterItem
// @Accept json
// @Produce json
// @Success 200 {object} model.CharacterItem "CharacterItem details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /CharacterItem [get]
func (a *CharacterItemApi) GetCharacterItems(ctx *gin.Context) {
	CharacterItems, err := a.DB.GetCharacterItems()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.CharacterItemExternal
	for _, CharacterItem := range CharacterItems {
		characterItemExternal := ToExternalCharacterItem(CharacterItem)
		resp = append(resp, characterItemExternal)
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateCharacterItem Updates CharacterItem by ID
//
// @Summary Updates CharacterItem by ID or nil
// @Description Permissions for Admin
// @Tags CharacterItem
// @Accept json
// @Produce json
// @Param id path int true "CharacterItem id"
// @Param characterItem body model.UpdateCharacterItem true "CharacterItem data"
// @Success 200 {object} model.CharacterItemExternal "CharacterItem details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "CharacterItem doesn't exist"
// @Router /CharacterItem/{id} [patch]
func (a *CharacterItemApi) UpdateCharacterItem(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var CharacterItem *model.UpdateCharacterItem
		if err := ctx.Bind(&CharacterItem); err == nil {
			oldCharacterItem, err := a.DB.GetCharacterItemByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldCharacterItem != nil {
				internal := &model.CharacterItem{
					ID:          oldCharacterItem.ID,
					CharacterId: CharacterItem.CharacterID,
					ItemId:      CharacterItem.ItemID,
					Quantity:    CharacterItem.Quantity,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateCharacterItem(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalCharacterItem(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "CharacterItem doesn't exist"})
		}
	})
}

// DeleteCharacterItem Deletes CharacterItem by ID
//
// @Summary Deletes CharacterItem by ID or returns nil
// @Description Permissions for Admin
// @Tags CharacterItem
// @Accept json
// @Produce json
// @Param id path int true "CharacterItem id"
// @Success 204
// @Failure 404 {string} string "CharacterItem doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /CharacterItem/{id} [delete]
func (a *CharacterItemApi) DeleteCharacterItem(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		CharacterItem, err := a.DB.GetCharacterItemByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if CharacterItem != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteCharacterItem(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalCharacterItem(characterItem *model.CharacterItem) *model.CharacterItemExternal {
	return &model.CharacterItemExternal{
		ID:          characterItem.ID,
		CharacterID: characterItem.CharacterId,
		ItemID:      characterItem.ItemId,
		Quantity:    characterItem.Quantity,
	}
}
