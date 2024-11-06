package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
	"strconv"
)

type SpellDatabase interface {
	CreateSpell(spell *model.Spell) error
	GetSpellByID(id uint) (*model.Spell, error)
	GetSpellByName(name string) (*model.Spell, error)
	GetSpells(limit int, offset int) ([]*model.Spell, error)
	DeleteSpell(id uint) error
	UpdateSpell(spell *model.Spell) error
	FindTraits(traitIDs []uint) ([]model.Trait, error)
	FindTraditions(IDs []uint) ([]model.Tradition, error)
}

type SpellAPI struct {
	DB SpellDatabase
}

// CreateSpell godoc
//
// @Summary Create and returns Spell or nil
// @Description Permissions for Admin
// @Tags Spell
// @Accept json
// @Produce json
// @Param Spell body model.SpellCreate true "spell data"
// @Success 201 {object} model.SpellExternal "spell details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /spell [post]
func (a *SpellAPI) CreateSpell(ctx *gin.Context) {
	spell := &model.SpellCreate{}
	if err := ctx.Bind(&spell); err == nil {
		traits, _ := a.DB.FindTraits(spell.TraitsID)
		traditions, _ := a.DB.FindTraditions(spell.TraditionID)
		internal := &model.Spell{
			Name:        spell.Name,
			Description: spell.Description,
			Rank:        spell.Rank,
			School:      &spell.School,
			Tradition:   traditions,
			Traits:      traits,
			Range:       spell.Range,
			Duration:    spell.Duration,
			Target:      spell.Target,
			Area:        spell.Area,
			Component:   spell.Component,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateSpell(internal)); !success {
			return
		}
	}
	ctx.JSON(http.StatusCreated, spell)
}

// GetSpells godoc
//
// @Summary Returns all Spells
// @Description Return all Spells
// @Tags Spell
// @Accept json
// @Produce json
// @Param limit query int false "Limit for pagination"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} model.SpellExternal "Spell details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /spell [get]
func (a *SpellAPI) GetSpells(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "0")
	offset := ctx.DefaultQuery("offset", "0")
	limitInt, err := strconv.Atoi(limit)
	offsetInt, err := strconv.Atoi(offset)
	spells, err := a.DB.GetSpells(limitInt, limitInt*offsetInt)
	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []model.SpellExternal
	for _, spell := range spells {
		externalSpell := ToExternalSpell(spell)
		resp = append(resp, *externalSpell)
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetSpellByID godoc
//
// @Summary Returns Spell by ID
// @Description Permissions for auth users
// @Tags Spell
// @Accept json
// @Produce json
// @Param id path int true "Spell id"
// @Success 200 {object} model.SpellExternal "Spell details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /spell/{id} [get]
func (a *SpellAPI) GetSpellByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		spell, err := a.DB.GetSpellByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, ToExternalSpell(spell))
	})
}

// UpdateSpell Updates Spell by ID
//
// @Summary Updates Spell by ID or nil
// @Description Permissions for Admin
// @Tags Spell
// @Accept json
// @Produce json
// @Param id path int true "Spell id"
// @Param Spell body model.SpellUpdate true "Spell data"
// @Success 200 {object} model.SpellExternal "Spell details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Spell doesn't exist"
// @Router /spell/{id} [patch]
func (a *SpellAPI) UpdateSpell(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var spell model.SpellUpdate
		if err := ctx.ShouldBindJSON(&spell); err == nil {
			oldSpell, err := a.DB.GetSpellByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			traits, _ := a.DB.FindTraits(spell.TraitsID)
			traditions, _ := a.DB.FindTraditions(spell.TraditionID)
			internalSpell := &model.Spell{
				ID:          oldSpell.ID,
				Name:        spell.Name,
				Description: spell.Description,
				Rank:        spell.Rank,
				School:      &spell.School,
				Tradition:   traditions,
				Traits:      traits,
				Range:       spell.Range,
				Duration:    spell.Duration,
				Target:      spell.Target,
				Area:        spell.Area,
				Component:   spell.Component,
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.UpdateSpell(internalSpell)); !success {
				ctx.JSON(http.StatusInternalServerError, success)
				return
			}
			newSpell, _ := a.DB.GetSpellByID(id)
			ctx.JSON(http.StatusOK, ToExternalSpell(newSpell))
		}
	})
}

// DeleteSpell Deletes Spell by ID
//
// @Summary Deletes Spell by ID or returns nil
// @Description Permissions for Admin
// @Tags Spell
// @Accept json
// @Produce json
// @Param id path int true "Spell id"
// @Success 204
// @Failure 404 {string} string "spell doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /spell/{id} [delete]
func (a *SpellAPI) DeleteSpell(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		spell, err := a.DB.GetSpellByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if spell != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteSpell(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalSpell(Spell *model.Spell) *model.SpellExternal {
	var traditon_names []string
	var trait_names []string
	for _, tradition_name := range Spell.Tradition {
		traditon_names = append(traditon_names, tradition_name.Name)
	}
	for _, trait_name := range Spell.Traits {
		trait_names = append(trait_names, trait_name.Name)
	}
	return &model.SpellExternal{
		ID:          Spell.ID,
		Name:        Spell.Name,
		Description: Spell.Description,
		Rank:        Spell.Rank,
		School:      *Spell.School,
		Tradition:   traditon_names,
		Traits:      trait_names,
		Range:       Spell.Range,
		Duration:    Spell.Duration,
		Target:      Spell.Target,
		Area:        Spell.Area,
		Component:   Spell.Component,
	}
}
