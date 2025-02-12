package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ClassFeatureDatabase interface {
	GetClassFeatureByID(id uint) (*model.ClassFeature, error)
	GetAllClassSkillFeatureByID(classFeatureID uint) ([]model.SkillFeature, error)
	GetClassFeatureByClassID(classID uint) ([]model.ClassFeature, error)
}

type ClassFeatureApi struct {
	DB ClassFeatureDatabase
}

// GetClassFeatureByID godoc
//
// @Summary Returns Class Feature by id
// @Description Retrieve Class Feature details using its ID
// @Tags Class Feature
// @Accept json
// @Produce json
// @Param id path int true "Class Feature id"
// @Success 200 {object} model.ClassFeatureExternal "Class Feature details"
// @Failure 404 {string} string "Class Feature not found"
// @Router /class-feature/{id} [get]
func (a *ClassFeatureApi) GetClassFeatureByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		classFeature, err := a.DB.GetClassFeatureByID(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Class Feature not found"})
			return
		}
		ctx.JSON(http.StatusOK, toExternalClassFeature(classFeature))
	})
}

// GetAllClassSkillFeatureByID godoc
//
// @Summary Returns Class Feature for certain level by id
// @Description Retrieve Class Feature for certain level details using its ID
// @Tags Class Feature
// @Accept json
// @Produce json
// @Param id path int true "Class Feature id"
// @Success 200 {object} model.SkillFeature "Class Feature details"
// @Failure 404 {string} string "Class Feature not found"
// @Router /class-feature/class/{id} [get]
func (a *ClassFeatureApi) GetAllClassSkillFeatureByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		classFeature, err := a.DB.GetAllClassSkillFeatureByID(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Class Feature not found"})
			return
		}
		ctx.JSON(http.StatusOK, classFeature)
	})
}

func toExternalClassFeature(classFeature *model.ClassFeature) *model.ClassFeatureExternal {
	var resp []*model.SkillFeatureExternal
	for _, feature := range classFeature.SkillFeatures {
		resp = append(resp, toExternalSkillFeature(feature))
	}

	return &model.ClassFeatureExternal{
		ID:                classFeature.ID,
		CharacterClassID:  classFeature.CharacterClassID,
		Level:             classFeature.Level,
		IsAncestryFeat:    classFeature.IsAncestryFeat,
		IsClassFeat:       classFeature.IsClassFeat,
		IsGeneralFeat:     classFeature.IsGeneralFeat,
		IsSkillFeat:       classFeature.IsSkillFeat,
		IsCharacterBoost:  classFeature.IsCharacterBoost,
		IsSkillIncrease:   classFeature.IsSkillIncrease,
		WeaponMastery:     classFeature.WeaponMastery,
		ArmorMastery:      classFeature.ArmorMastery,
		PerceptionMastery: classFeature.PerceptionMastery,
		FortitudeMastery:  classFeature.FortitudeMastery,
		ReflexMastery:     classFeature.ReflexMastery,
		WillMastery:       classFeature.WillMastery,
		SkillFeature:      resp,
	}
}

func toExternalSkillFeature(classSkillFeature model.SkillFeature) *model.SkillFeatureExternal {
	return &model.SkillFeatureExternal{
		ID:          classSkillFeature.ID,
		Name:        classSkillFeature.Name,
		Description: classSkillFeature.Description,
	}
}
