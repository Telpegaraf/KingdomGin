package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

// CreateArmor godoc
//
// @Summary Create and returns armor or nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param character body model.CreateArmor true "Armor data"
// @Success 201 {object} model.ArmorExternal "Armor details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /item/armor [post]
func (a *ItemApi) CreateArmor(ctx *gin.Context) {
	armor := &model.CreateArmor{}
	if err := ctx.ShouldBindJSON(armor); err == nil {
		internalArmor := &model.Armor{
			ArmorClass: armor.ArmorClass,
		}
		internalItem := &model.Item{
			Name:        armor.Name,
			Description: armor.Description,
			Bulk:        armor.Bulk,
			Level:       armor.Level,
			Price:       armor.Price,
			OwnerType:   "armors",
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateArmor(internalArmor, internalItem)); !success {
			ctx.JSON(http.StatusInternalServerError, success)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalArmor(internalArmor, internalItem))
	}
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

// GetArmorByID godoc
//
// @Summary Returns Armor by ID
// @Description Permissions for auth users
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "armor id"
// @Success 200 {object} model.Armor "Armor details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/armor/{id} [get]
func (a *ItemApi) GetArmorByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		armor, err := a.DB.GetArmorByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalArmor(armor, &armor.Item))
	})
}

// UpdateArmor Updates Armor by ID
//
// @Summary Updates Armor by ID or nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Armor id"
// @Param character body model.UpdateArmor true "Armor data"
// @Success 200 {object} model.Armor "Armor details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Domain doesn't exist"
// @Router /item/armor/{id} [patch]
func (a *ItemApi) UpdateArmor(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var armor model.UpdateArmor
		if err := ctx.ShouldBindJSON(&armor); err == nil {
			oldArmor, err := a.DB.GetArmorByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			internalArmor := &model.Armor{
				ArmorClass: armor.ArmorClass,
				ID:         oldArmor.ID,
			}
			internalItem := &model.Item{
				ID:          oldArmor.Item.ID,
				Name:        armor.Name,
				Description: armor.Description,
				Bulk:        armor.Bulk,
				Level:       armor.Level,
				Price:       armor.Price,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateArmor(internalArmor, internalItem)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			ctx.JSON(http.StatusOK, ToExternalArmor(internalArmor, internalItem))
		}
	})
}

func ToExternalArmor(armor *model.Armor, item *model.Item) *model.ArmorExternal {
	return &model.ArmorExternal{
		ID:          armor.ID,
		Name:        item.Name,
		Description: item.Description,
		Level:       item.Level,
		Bulk:        item.Bulk,
		Price:       item.Price,
		ArmorClass:  armor.ArmorClass,
		ItemID:      item.ID,
	}
}
