package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

func (a *CharacterApi) CreateCharacterFeat(characterID uint, background *model.Background) {
	characterFeat := &model.CharacterFeat{
		CharacterID: characterID,
		FeatID:      background.FeatID,
	}
	err := a.DB.CreateCharacterFeat(characterFeat)
	if err != nil {
		return
	}
}

func (a *CharacterApi) AddCharacterFeat(ctx *gin.Context) {
	characterFeat := &model.CreateCharacterFeat{}
	if err := ctx.ShouldBindJSON(characterFeat); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feat, err := a.DB.GetFeatByID(characterFeat.FeatID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if feat.PrerequisiteSkillID != nil {
		characterSkill, err := a.DB.GetCharacterSkillByID(characterFeat.CharacterID, *feat.PrerequisiteSkillID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !a.checkSkillMastery(characterSkill.Mastery, feat.PrerequisiteMastery) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong Skill Mastery"})
			return
		}
	}

	internal := model.CharacterFeat{
		CharacterID: characterFeat.CharacterID,
		FeatID:      characterFeat.FeatID,
	}
	if err := a.DB.CreateCharacterFeat(&internal); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, internal)
}

func (a *CharacterApi) checkSkillMastery(actual, required model.MasteryLevel) bool {
	masteryOrder := map[model.MasteryLevel]int{
		model.None:   0,
		model.Train:  1,
		model.Expert: 2,
		model.Master: 3,
		model.Legend: 4,
	}

	return masteryOrder[actual] >= masteryOrder[required]
}
