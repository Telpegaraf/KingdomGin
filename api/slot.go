package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type SlotDatabase interface {
	GetSlotByID(id uint) (*model.Slot, error)
	UpdateSlot(slot *model.Slot) error
}

type SlotApi struct {
	DB SlotDatabase
}

// CreateSlot creates slot linked with Character
func (a *CharacterApi) CreateSlot(ctx *gin.Context, characterID uint) {
	internal := &model.Slot{
		CharacterID: characterID,
	}
	if success := SuccessOrAbort(ctx, 500, a.DB.CreateSlot(internal)); !success {
		return
	}
}

// GetSlotByID godoc
//
// @Summary Returns slot by id
// @Description Permissions for auth user
// @Tags Slot
// @Accept json
// @Produce json
// @Param id path int true "slot id"
// @Success 200 {object} model.SlotExternal "slot details"
// @Failure 404 {string} string "Slot not found"
// @Router /slot/{id} [get]
func (a *SlotApi) GetSlotByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		slot, err := a.DB.GetSlotByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if slot != nil {
			ctx.JSON(http.StatusOK, ToExternalSlot(slot))
		}
	})
}

// UpdateSlot Updates Slot by ID
//
// @Summary Updates Slot by ID or nil
// @Description Permissions for Character's User or Admin
// @Tags Slot
// @Accept json
// @Produce json
// @Param id path int true "Slot id"
// @Param slot body model.SlotUpdate true "Character data"
// @Success 200 {object} model.SlotExternal "Character details"
// @Failure 404 {string} string "Slot doesn't exist"
// @Router /slot/{id} [patch]
func (a *SlotApi) UpdateSlot(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var slot *model.Slot
		if err := ctx.ShouldBindJSON(&slot); err == nil {
			oldSlot, err := a.DB.GetSlotByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldSlot != nil {
				internal := &model.Slot{
					ID:             oldSlot.ID,
					CharacterID:    oldSlot.CharacterID,
					ArmorID:        slot.ArmorID,
					FirstWeaponID:  slot.FirstWeaponID,
					SecondWeaponID: slot.SecondWeaponID,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateSlot(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalSlot(oldSlot))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Slot doesn't exist"})
		}
	})
}

func ToExternalSlot(slot *model.Slot) *model.SlotExternal {
	return &model.SlotExternal{
		ID:             slot.ID,
		CharacterID:    slot.CharacterID,
		ArmorID:        slot.ArmorID,
		FirstWeaponID:  slot.FirstWeaponID,
		SecondWeaponID: slot.SecondWeaponID,
	}
}
