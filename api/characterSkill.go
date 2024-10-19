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

func (a *CharacterSkillApi) CharacterSkillCreate(ctx *gin.Context) {
	characterSkill := &model.CharacterSkillCreate{}
	if err := ctx.ShouldBindJSON(characterSkill); err == nil {
		internal := &model.CharacterSkill{
			CharacterID: characterSkill.CharacterID,
			SkillID:     characterSkill.SkillID,
			Mastery:     characterSkill.Mastery,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CharacterSkillCreate(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToExternalCharacterSkill(internal))
	}
}

func (a *CharacterSkillApi) GetCharacterSkillByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		characterSkill, err := a.DB.GetCharacterSkillByID(id)
		if success := SuccessOrAbort(ctx, 404, err); success {
			return
		}
		if characterSkill != nil {
			ctx.JSON(http.StatusOK, ToExternalCharacterSkill(characterSkill))
		}
	})
}

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
					SkillID:     oldCharacterSkill.SkillID,
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
		SkillID:     characterSKill.SkillID,
		Mastery:     characterSKill.Mastery,
	}
}
