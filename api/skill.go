package api

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"kingdom/model"
	"log"
	"net/http"
	"os"
)

type SkillDatabase interface {
	GetSkillByID(id uint) (*model.Skill, error)
	GetSkillByName(name string) (*model.Skill, error)
	CreateSkill(Skill *model.Skill) error
	GetSkills() ([]*model.Skill, error)
	UpdateSkill(Skill *model.Skill) error
	DeleteSkill(id uint) error
}

type SkillApi struct {
	DB SkillDatabase
}

// CreateSkill godoc
//
// @Summary Create and returns Skill or nil
// @Description Permissions for Admin
// @Tags Skill
// @Accept json
// @Produce json
// @Param Skill body model.SkillCreate true "Skill data"
// @Success 201 {object} model.SkillExternal "Skill details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /skill [post]
func (a *SkillApi) CreateSkill(ctx *gin.Context) {
	skill := &model.SkillCreate{}
	if err := ctx.ShouldBindJSON(skill); err == nil {
		internal := &model.Skill{
			Name:        skill.Name,
			Description: skill.Description,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateSkill(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToSkillExternal(internal))
	}
}

// GetSkillByID godoc
//
// @Summary Returns Skill by id
// @Description Retrieve Skill details using its ID
// @Tags Skill
// @Accept json
// @Produce json
// @Param id path int true "Skill id"
// @Success 200 {object} model.SkillExternal "Skill details"
// @Failure 404 {string} string "Skill not found"
// @Router /skill/{id} [get]
func (a *SkillApi) GetSkillByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		skill, err := a.DB.GetSkillByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if skill != nil {
			ctx.JSON(http.StatusOK, ToSkillExternal(skill))
		}
	})
}

// LoadSkill godoc
//
// @Summary Create Skill from csv file on server or nil
// @Description Permissions for Admin
// @Tags Skill
// @Accept json
// @Produce json
// @Success 201 {object} model.SkillExternal "Skill details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /skill/load [post]
func (a *SkillApi) LoadSkill(ctx *gin.Context) {
	file, err := os.Open("./csv/Skill.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := csv.NewReader(file)
	reader.Comma = ';'
	var Skills []model.Skill

	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Skill := model.Skill{
			Name:        record[0],
			Description: record[1],
			Ability:     model.Ability(record[2]),
		}
		Skills = append(Skills, Skill)
		if existSkill, err := a.DB.GetSkillByName(Skill.Name); err == nil && existSkill != nil {
			continue
		}
		err = a.DB.CreateSkill(&Skill)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

// GetSkills godoc
//
// @Summary Returns all Skills
// @Description Return all Skills
// @Tags Skill
// @Accept json
// @Produce json
// @Success 200 {object} model.SkillExternal "Skill details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /skill [get]
func (a *SkillApi) GetSkills(ctx *gin.Context) {
	skills, err := a.DB.GetSkills()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.SkillExternal
	for _, skill := range skills {
		resp = append(resp, ToSkillExternal(skill))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateSkill Updates Skill by ID
//
// @Summary Updates Skill by ID or nil
// @Description Permissions for Admin
// @Tags Skill
// @Accept json
// @Produce json
// @Param id path int true "Skill id"
// @Param character body model.SkillUpdate true "Skill data"
// @Success 200 {object} model.SkillExternal "Skill details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Skill doesn't exist"
// @Router /skill/{id} [patch]
func (a *SkillApi) UpdateSkill(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var Skill *model.SkillUpdate
		if err := ctx.Bind(&Skill); err == nil {
			oldSkill, err := a.DB.GetSkillByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldSkill != nil {
				internal := &model.Skill{
					ID:          oldSkill.ID,
					Name:        Skill.Name,
					Description: Skill.Description,
					Ability:     Skill.Ability,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateSkill(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToSkillExternal(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Skill doesn't exist"})
		}
	})
}

// DeleteSkill Deletes Skill by ID
//
// @Summary Deletes Skill by ID or returns nil
// @Description Permissions for Admin
// @Tags Skill
// @Accept json
// @Produce json
// @Param id path int true "Skill id"
// @Success 204
// @Failure 404 {string} string "Skill doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /skill/{id} [delete]
func (a *SkillApi) DeleteSkill(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		skill, err := a.DB.GetSkillByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if skill != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteSkill(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Skill was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Skill doesn't exist"})
		}
	})
}

func ToSkillExternal(Skill *model.Skill) *model.SkillExternal {
	return &model.SkillExternal{
		ID:          Skill.ID,
		Name:        Skill.Name,
		Description: Skill.Description,
		Ability:     Skill.Ability,
	}
}
