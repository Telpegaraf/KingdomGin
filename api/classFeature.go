package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ClassFeatureDatabase interface {
	GetClassFeatureByID(id uint) (*model.ClassFeature, error)
	GetSkillFeatureByID(classFeatureID uint) ([]model.SkillFeature, error)
	GetClassFeatureByClassID(classID uint) ([]model.ClassFeature, error)
}

type ClassFeatureApi struct {
	DB ClassFeatureDatabase
}

// GetAllFeature godoc
//
// @Summary Returns All Class Feature by id
// @Description Retrieve All Class Feature details using its ID
// @Tags Class Feature
// @Accept json
// @Produce json
// @Param id path int true "Class Feature id"
// @Success 200 {object} model.ClassFeatureExternal "All Class Feature details"
// @Failure 404 {string} string "Class Feature not found"
// @Router /class-feature/all/{id} [get]
func (a *ClassFeatureApi) GetAllFeature(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		features, err := a.DB.GetClassFeatureByClassID(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "features for this class not found"})
			return
		}
		var resp []*model.ClassFeatureExternal
		for _, feature := range features {
			resp = append(resp, toExternalClassFeature(&feature))
		}
		ctx.JSON(http.StatusOK, gin.H{"data": resp})
	})
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

// GetClassSkillFeatureByID godoc
//
// @Summary Returns Class Feature for certain level by id
// @Description Retrieve Class Feature for certain level details using its ID
// @Tags Class Feature
// @Accept json
// @Produce json
// @Param id path int true "Class Feature id"
// @Success 200 {object} model.SkillFeature "Class Feature details"
// @Failure 404 {string} string "Class Feature not found"
// @Router /skill-feature/{id} [get]
func (a *ClassFeatureApi) GetClassSkillFeatureByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		skillFeatures, err := a.DB.GetSkillFeatureByID(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Class Skill Feature not found"})
			return
		}
		var resp []*model.SkillFeatureExternal
		for _, feature := range skillFeatures {
			resp = append(resp, toExternalSkillFeature(feature))
		}
		ctx.JSON(http.StatusOK, resp)
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
