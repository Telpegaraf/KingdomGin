package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ItemDatabase interface {
	GetItems() ([]*model.Item, error)
	GetArmors() ([]*model.Armor, error)
	GetWeapons() ([]*model.Weapon, error)
	GetGears() ([]*model.Gear, error)
	GetArmorByID(id uint) (*model.Armor, error)
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

// GetWeapons godoc
//
// @Summary Returns all weapons
// @Description Return all weapons
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} model.Weapon "Weapon details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/weapon [get]
func (a *ItemApi) GetWeapons(ctx *gin.Context) {
	weapons, err := a.DB.GetWeapons()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.Weapon
	for _, weapon := range weapons {
		resp = append(resp, weapon)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetGears godoc
//
// @Summary Returns all gears
// @Description Return all gears
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} model.Gear "Gear details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/gear [get]
func (a *ItemApi) GetGears(ctx *gin.Context) {
	gears, err := a.DB.GetGears()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.Gear
	for _, gear := range gears {
		resp = append(resp, gear)
	}
	ctx.JSON(http.StatusOK, resp)
}
