package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type SlotDatabase interface {
	GetSlotByID(id uint) (*model.Slot, error)
	UpdateSlot(slot *model.Slot) error
	GetCharacterItems(characterId uint) ([]*model.CharacterItem, error)
	GetCharacterItemByID(id uint) (*model.CharacterItem, error)
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
			//ctx.JSON(http.StatusOK, slot)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Slot doesn't exist"})
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
		var slot *model.SlotUpdate
		if err := ctx.ShouldBindJSON(&slot); err == nil {

			if slot.ArmorID != nil {
				if armor, err := a.DB.GetCharacterItemByID(*slot.ArmorID); err == nil && armor != nil &&
					armor.Item.OwnerType != "armors" {
					ctx.JSON(http.StatusNotFound, gin.H{"error": "Wrong Armor slot"})
					return
				}
			}

			if slot.FirstWeaponID != nil {
				if fWeapon, err := a.DB.GetCharacterItemByID(*slot.FirstWeaponID); err == nil && fWeapon != nil &&
					fWeapon.Item.OwnerType != "weapons" {
					ctx.JSON(http.StatusNotFound, gin.H{"error": "Wrong First Weapon slot"})
					return
				}
			}
			if slot.SecondWeaponID != nil {
				if sWeapon, err := a.DB.GetCharacterItemByID(*slot.SecondWeaponID); err == nil && sWeapon != nil &&
					sWeapon.Item.OwnerType != "weapons" {
					ctx.JSON(http.StatusNotFound, gin.H{"error": "Wrong Second Weapon slot"})
					return
				}
			}

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
				ctx.JSON(http.StatusOK, ToExternalSlot(internal))
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Slot doesn't exist"})
			}
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

//func CheckSlot(
//	characterItems []*model.CharacterItem,
//	armorId *uint,
//	firstWeaponId *uint,
//	secondWeaponId *uint) bool {
//	var armorValue uint
//	var firstWeaponValue uint
//	var secondWeaponValue uint
//	if armorId != nil {
//		armorValue = *armorId
//	}
//	if firstWeaponId != nil {
//		firstWeaponValue = *firstWeaponId
//	}
//	if secondWeaponId != nil {
//		secondWeaponValue = *secondWeaponId
//	}
//	for _, characterItem := range characterItems {
//		if characterItem.ID == armorValue {
//			if characterItem.Item.OwnerType != "armors" {
//				return false
//			}
//		}
//		if characterItem.ID == firstWeaponValue {
//			if characterItem.Item.OwnerType != "weapons" {
//				return false
//			}
//		}
//		if characterItem.ID == secondWeaponValue {
//			if characterItem.Item.OwnerType != "weapons" {
//				return false
//			}
//		}
//	}
//	return true
//}
