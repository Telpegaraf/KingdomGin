package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ItemDatabase interface {
	GetItems() ([]*model.Item, error)
	GetItemByID(id uint) (*model.Item, error)
	GetArmors() ([]*model.Armor, error)
	GetArmorByID(id uint) (*model.Armor, error)
	CreateArmor(armor *model.Armor, item *model.Item) error
	UpdateArmor(armor *model.Armor, item *model.Item) error
	GetWeapons() ([]*model.Weapon, error)
	GetWeaponByID(id uint) (*model.Weapon, error)
	CreateWeapon(weapon *model.Weapon, item *model.Item) error
	UpdateWeapon(weapon *model.Weapon, item *model.Item) error
	GetGears() ([]*model.Gear, error)
	GetGearByID(id uint) (*model.Gear, error)
	CreateGear(weapon *model.Gear, item *model.Item) error
	UpdateGear(weapon *model.Gear, item *model.Item) error
	DeleteItem(id uint, ownerType string, ownerID uint) error
}

type ItemApi struct {
	DB ItemDatabase
}

// GetItems godoc
//
// @Summary Returns all items
// @Description Return all items
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} model.Item "Item details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item [get]
func (a *ItemApi) GetItems(ctx *gin.Context) {
	items, err := a.DB.GetItems()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.Item
	for _, item := range items {
		resp = append(resp, item)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetItemByID godoc
//
// @Summary Returns Item by ID
// @Description Permissions for auth users
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item id"
// @Success 200 {object} model.ItemExternal "Item details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/{id} [get]
func (a *ItemApi) GetItemByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		item, err := a.DB.GetItemByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToItemExternal(item))
	})
}

// DeleteItem Deletes Item by ID
//
// @Summary Deletes Item by ID or returns nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item id"
// @Success 204
// @Failure 404 {string} string "Item doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /item/{id} [delete]
func (a *ItemApi) DeleteItem(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		item, err := a.DB.GetItemByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if item != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteItem(id, item.OwnerType, item.OwnerID)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToItemExternal(item *model.Item) *model.ItemExternal {
	return &model.ItemExternal{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Level:       item.Level,
		Bulk:        item.Bulk,
		Price:       item.Price,
		OwnerType:   item.OwnerType,
		OwnerID:     item.OwnerID,
	}
}
