package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

// CreateWeapon godoc
//
// @Summary Create and returns weapon or nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param weapon body model.CreateWeapon true "Weapon data"
// @Success 201 {object} model.WeaponExternal "Weapon details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /item/weapon [post]
func (a *ItemApi) CreateWeapon(ctx *gin.Context) {
	weapon := &model.CreateWeapon{}
	if err := ctx.ShouldBindJSON(weapon); err == nil {
		internalWeapon := &model.Weapon{
			Dice:         weapon.Dice,
			DiceQuantity: weapon.DiceQuantity,
			Damage:       *weapon.Damage,
			DamageType:   weapon.DamageType,
		}
		internalItem := &model.Item{
			Name:        weapon.Name,
			Description: weapon.Description,
			Bulk:        weapon.Bulk,
			Level:       *weapon.Level,
			Price:       weapon.Price,
			OwnerType:   "weapons",
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateWeapon(internalWeapon, internalItem)); !success {
			ctx.JSON(http.StatusInternalServerError, success)
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalWeapon(internalWeapon, internalItem))
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
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
	var resp []*model.WeaponExternal
	for _, weapon := range weapons {
		externalWeapon := ToExternalWeapon(weapon, &weapon.Item)
		resp = append(resp, externalWeapon)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetWeaponByID godoc
//
// @Summary Returns weapon by ID
// @Description Permissions for auth users
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Weapon id"
// @Success 200 {object} model.Weapon "Weapon details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /item/weapon/{id} [get]
func (a *ItemApi) GetWeaponByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		weapon, err := a.DB.GetWeaponByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalWeapon(weapon, &weapon.Item))
	})
}

// UpdateWeapon Updates weapon by ID
//
// @Summary Updates weapon by ID or nil
// @Description Permissions for Admin
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Weapon id"
// @Param character body model.UpdateWeapon true "Weapon data"
// @Success 200 {object} model.Weapon "Weapon details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Weapon doesn't exist"
// @Router /item/weapon/{id} [patch]
func (a *ItemApi) UpdateWeapon(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var weapon model.UpdateWeapon
		if err := ctx.ShouldBindJSON(&weapon); err == nil {
			oldWeapon, err := a.DB.GetWeaponByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			internalWeapon := &model.Weapon{
				Dice:         weapon.Dice,
				DiceQuantity: weapon.DiceQuantity,
				Damage:       *weapon.Damage,
				DamageType:   weapon.DamageType,
				ID:           oldWeapon.ID,
			}
			internalItem := &model.Item{
				ID:          oldWeapon.Item.ID,
				Name:        weapon.Name,
				Description: weapon.Description,
				Bulk:        weapon.Bulk,
				Level:       *weapon.Level,
				Price:       weapon.Price,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateWeapon(internalWeapon, internalItem)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			newWeapon, _ := a.DB.GetWeaponByID(id)
			ctx.JSON(http.StatusOK, ToExternalWeapon(newWeapon, &newWeapon.Item))
		}
	})
}

func ToExternalWeapon(weapon *model.Weapon, item *model.Item) *model.WeaponExternal {
	return &model.WeaponExternal{
		ID:           weapon.ID,
		Name:         item.Name,
		Description:  item.Description,
		Level:        item.Level,
		Bulk:         item.Bulk,
		Price:        item.Price,
		Dice:         weapon.Dice,
		DiceQuantity: weapon.DiceQuantity,
		Damage:       weapon.Damage,
		DamageType:   weapon.DamageType,
		ItemID:       item.ID,
	}
}
