package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ItemDatabase interface {
	GetItems() ([]*model.Item, error)
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
