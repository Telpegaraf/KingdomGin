package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterSkillDatabase interface {
	CharacterSkillCreate(characterSkill *model.CharacterSkill) error
	GetCharacterSkillByID(id uint) (*model.CharacterSkill, error)
	GetCharacterSkills(id uint) ([]*model.CharacterSkill, error)
	UpdateCharacterSkill(skill *model.CharacterSkill) error
}

type CharacterSkillApi struct {
	DB CharacterSkillDatabase
}

// CharacterSkillCreate godoc
//
// @Summary Create and returns Character Skill or nil
// @Description Permissions for Admin
// @Tags Character Skill
// @Accept json
// @Produce json
// @Param characterSkill body model.CharacterSkillCreate true "Action data"
// @Success 201 {object} model.CharacterSkillExternal "Action details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /character-skill [post]
func (a *CharacterSkillApi) CharacterSkillCreate(ctx *gin.Context) {
	characterSkill := &model.CharacterSkillCreate{}
	if err := ctx.ShouldBindJSON(characterSkill); err == nil {
		internal := &model.CharacterSkill{
			CharacterID: characterSkill.CharacterID,
			Name:        characterSkill.Name,
			Mastery:     characterSkill.Mastery,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CharacterSkillCreate(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalCharacterSkill(internal))
	}
}

// GetCharacterSkills godoc
//
// @Summary Returns all Character Skills
// @Description Return all Character Skills
// @Tags Character Skill
// @Accept json
// @Produce json
// @Success 200 {object} model.CharacterSkill "Character Skill details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /character-skill [get]
func (a *CharacterSkillApi) GetCharacterSkills(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		characterSkills, err := a.DB.GetCharacterSkills(id)
		if success := SuccessOrAbort(ctx, 500, err); success {
			return
		}
		var resp []*model.CharacterSkillExternal
		for _, characterSkill := range characterSkills {
			resp = append(resp, ToExternalCharacterSkill(characterSkill))
		}
		ctx.JSON(http.StatusOK, resp)
	})
}

// UpdateCharacterSkill Updates Character Skill by ID
//
// @Summary Updates Character Skill by ID or nil
// @Description Permissions for Admin
// @Tags Character Skill
// @Accept json
// @Produce json
// @Param id path int true "Character Skill id"
// @Param characterSkill body model.CharacterSkillUpdate true "Character Skill data"
// @Success 200 {object} model.CharacterSkillExternal "Action details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Character Skill doesn't exist"
// @Router /character-skill/{id} [patch]
func (a *CharacterSkillApi) UpdateCharacterSkill(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var characterSkill *model.CharacterSkillUpdate
		if err := ctx.ShouldBindJSON(&characterSkill); err == nil {
			oldCharacterSkill, err := a.DB.GetCharacterSkillByID(id)
			if success := SuccessOrAbort(ctx, 500, err); success {
				return
			}
			if oldCharacterSkill != nil {
				internal := &model.CharacterSkill{
					ID:          oldCharacterSkill.ID,
					CharacterID: oldCharacterSkill.CharacterID,
					Name:        oldCharacterSkill.Name,
					Mastery:     characterSkill.Mastery,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateCharacterSkill(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalCharacterSkill(internal))
			}
		}
	})
}

func ToExternalCharacterSkill(characterSKill *model.CharacterSkill) *model.CharacterSkillExternal {
	return &model.CharacterSkillExternal{
		ID:          characterSKill.ID,
		CharacterID: characterSKill.CharacterID,
		Name:        characterSKill.Name,
		Mastery:     characterSKill.Mastery,
	}
}
