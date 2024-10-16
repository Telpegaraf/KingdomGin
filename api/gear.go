package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

// CreateGear godoc
//
// @Summary Create and returns Gear or nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param gear body model.CreateGear true "Gear data"
// @Success 201 {object} model.GearExternal "Gear details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /item/gear [post]
func (a *ItemApi) CreateGear(ctx *gin.Context) {
	Gear := &model.CreateGear{}
	if err := ctx.ShouldBindJSON(Gear); err == nil {
		internalGear := &model.Gear{}
		internalItem := &model.Item{
			Name:        Gear.Name,
			Description: Gear.Description,
			Bulk:        Gear.Bulk,
			Level:       *Gear.Level,
			Price:       Gear.Price,
			OwnerType:   "gears",
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateGear(internalGear, internalItem)); !success {
			ctx.JSON(http.StatusInternalServerError, success)
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalGear(internalGear, internalItem))
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
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
	var resp []*model.GearExternal
	for _, gear := range gears {
		externalGear := ToExternalGear(gear, &gear.Item)
		resp = append(resp, externalGear)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetGearByID godoc
//
// @Summary Returns Gear by ID
// @Description Permissions for auth users
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Gear id"
// @Success 200 {object} model.Gear "Gear details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/gear/{id} [get]
func (a *ItemApi) GetGearByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		Gear, err := a.DB.GetGearByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalGear(Gear, &Gear.Item))
	})
}

// UpdateGear Updates Gear by ID
//
// @Summary Updates Gear by ID or nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Gear id"
// @Param Gear body model.UpdateGear true "Gear data"
// @Success 200 {object} model.Gear "Gear details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Gear doesn't exist"
// @Router /item/gear/{id} [patch]
func (a *ItemApi) UpdateGear(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var Gear model.UpdateGear
		if err := ctx.ShouldBindJSON(&Gear); err == nil {
			oldGear, err := a.DB.GetGearByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			internalGear := &model.Gear{
				ID: oldGear.ID,
			}
			internalItem := &model.Item{
				ID:          oldGear.Item.ID,
				Name:        Gear.Name,
				Description: Gear.Description,
				Bulk:        Gear.Bulk,
				Level:       *Gear.Level,
				Price:       Gear.Price,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateGear(internalGear, internalItem)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			newGear, _ := a.DB.GetGearByID(id)
			ctx.JSON(http.StatusOK, ToExternalGear(newGear, &newGear.Item))
		}
	})
}

func ToExternalGear(Gear *model.Gear, item *model.Item) *model.GearExternal {
	return &model.GearExternal{
		ID:          Gear.ID,
		Name:        item.Name,
		Description: item.Description,
		Level:       item.Level,
		Bulk:        item.Bulk,
		Price:       item.Price,
		ItemID:      item.ID,
	}
}
