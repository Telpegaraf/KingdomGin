package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
)

type CharacterClassDatabase interface {
	CreateCharacterClass(characterClass *model.CharacterClass) error
	GetCharacterClasses() (*[]model.CharacterClass, error)
	DeleteCharacterClass(id uint) error
	UpdateCharacterClass(class *model.CharacterClass) error
	GetTraditionByName(name string) (*model.Tradition, error)
}

type CharacterClassApi struct {
	DB CharacterClassDatabase
}

// CreateCharacterClass godoc
//
// @Summary Create and returns new Character Class or nil
// @Description Create new Character Class
// @Tags Character Class
// @Accept json
// @Produce json
// @Param god body model.CharacterClassCreate true "Character Class data"
// @Success 201 {object} model.CharacterClass "Character Class details"
// @Failure 401 {string} string "Unauthorized"
// @Router /class [post]
func (a *CharacterClassApi) CreateCharacterClass(ctx *gin.Context) {
	characterClass := &model.CharacterClassCreate{}
	if err := ctx.Bind(characterClass); err == nil {
		internal := &model.CharacterClass{
			Name:     characterClass.Name,
			HitPoint: characterClass.HitPoint,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacterClass(internal)); !success {
			return
		}
	}
}
