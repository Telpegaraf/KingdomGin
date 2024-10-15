package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ItemDatabase interface {
	GetItems() ([]*model.Item, error)
	GetArmors() ([]*model.Armor, error)
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
// @Success 200 {object} model.ItemExternal "Item details"
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

// GetArmors godoc
//
// @Summary Returns all armors
// @Description Return all armors
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} model.Armor "Armor details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/armor [get]
func (a *ItemApi) GetArmors(ctx *gin.Context) {
	armors, err := a.DB.GetArmors()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.Armor
	for _, armor := range armors {
		resp = append(resp, armor)
	}
	ctx.JSON(http.StatusOK, resp)
}
