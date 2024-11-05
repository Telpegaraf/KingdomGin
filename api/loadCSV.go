package api

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"kingdom/model"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type LoadCSVDatabase interface {
	GetTraditionByName(name string) (*model.Tradition, error)
	CreateTradition(Tradition *model.Tradition) error
	GetSkillByName(name string) (*model.Skill, error)
	CreateSkill(Skill *model.Skill) error
	GetActionByName(name string) (*model.Action, error)
	CreateAction(Action *model.Action) error
	CreateCharacterClass(characterClass *model.CharacterClass) error
	GetCharacterClassByName(name string) (*model.CharacterClass, error)
	GetTraitByName(name string) (*model.Trait, error)
	CreateTrait(Trait *model.Trait) error
	GetFeatByName(name string) (*model.Feat, error)
	CreateFeat(feat *model.Feat) error
	FindTraits(traitIDs []uint) ([]model.Trait, error)
	FindTraditions(IDs []uint) ([]model.Tradition, error)
	CreateSpell(spell *model.Spell) error
	GetDomainByName(name string) (*model.Domain, error)
}

type LoadCSVApi struct {
	DB LoadCSVDatabase
}

// LoadCSV godoc
//
// @Summary Create and returns models from csv files or nil
// @Description Permissions for Admin, csv - Tradition, Character Class, Trait, Action, Skill, Feat
// @Tags CSV
// @Accept json
// @Produce json
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /csv [post]
func (a *LoadCSVApi) LoadCSV(ctx *gin.Context) {
	a.LoadTradition(ctx)
	a.LoadCharacterClass(ctx)
	a.LoadTrait(ctx)
	a.LoadAction(ctx)
	a.LoadSkill(ctx)
	a.LoadFeat(ctx)
	a.LoadSpell(ctx)
}

func (a *LoadCSVApi) LoadTradition(ctx *gin.Context) {
	file, err := os.Open("./csv/Tradition.csv")
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
	var traditions []model.Tradition

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tradition := model.Tradition{
			Name:        record[0],
			Description: record[1],
		}
		traditions = append(traditions, tradition)
		if existTradition, err := a.DB.GetTraditionByName(tradition.Name); err == nil && existTradition != nil {
			continue
		}
		err = a.DB.CreateTradition(&tradition)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *LoadCSVApi) LoadTrait(ctx *gin.Context) {
	file, err := os.Open("./csv/Trait.csv")
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
	var traits []model.Trait

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
		trait := model.Trait{
			Name:        record[0],
			Description: record[1],
		}
		traits = append(traits, trait)
		if existTrait, err := a.DB.GetTraitByName(trait.Name); err == nil && existTrait != nil {
			continue
		}
		err = a.DB.CreateTrait(&trait)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *LoadCSVApi) LoadSkill(ctx *gin.Context) {
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

func (a *LoadCSVApi) LoadAction(ctx *gin.Context) {
	file, err := os.Open("./csv/Action.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := csv.NewReader(file)
	var actions []model.Action

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		action := model.Action{
			Name: record[0],
		}
		actions = append(actions, action)
		if existAction, err := a.DB.GetActionByName(action.Name); err == nil && existAction != nil {
			continue
		}
		err = a.DB.CreateAction(&action)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *LoadCSVApi) LoadSpell(ctx *gin.Context) {
	file, err := os.Open("./csv/Spell.csv")
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
	var spells []model.Spell

	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(len(record))
			log.Fatal(err)
		}

		log.Println(record)

		rank, err := strconv.Atoi(record[11])
		traits := a.GetTraits(ctx, record[10])
		traditions := a.GetTraditions(ctx, record[9])

		spell := model.Spell{
			Name:        record[0],
			Description: record[1],
			Component:   record[2],
			Range:       record[3],
			Area:        record[4],
			Duration:    record[5],
			Target:      record[6],
			Cast:        record[8],
			Rank:        uint8(rank),
		}

		if record[7] == "" {
			continue
		}

		if traits != nil {
			spellTraits, err := a.DB.FindTraits(traits)
			if err != nil {
				log.Fatal(err)
			}
			spell.Traits = spellTraits
		}

		if traditions != nil {
			spellTraditions, err := a.DB.FindTraditions(traditions)
			if err != nil {
				log.Fatal(err)
			}
			spell.Traditional = spellTraditions
		}

		spells = append(spells, spell)
		err = a.DB.CreateSpell(&spell)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *LoadCSVApi) LoadCharacterClass(ctx *gin.Context) {
	file, err := os.Open("./csv/CharacterClass.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := csv.NewReader(file)
	var characterClasses []model.CharacterClass

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		hitPoint, _ := strconv.Atoi(record[1])

		characterClass := model.CharacterClass{
			Name:          record[0],
			HitPoint:      uint16(hitPoint),
			Perception:    model.MasteryLevel(record[2]),
			Fortitude:     model.MasteryLevel(record[3]),
			Reflex:        model.MasteryLevel(record[4]),
			Will:          model.MasteryLevel(record[5]),
			UnarmedArmor:  model.MasteryLevel(record[6]),
			LightArmor:    model.MasteryLevel(record[7]),
			MediumArmor:   model.MasteryLevel(record[8]),
			HeavyArmor:    model.MasteryLevel(record[9]),
			UnArmedWeapon: model.MasteryLevel(record[10]),
			CommonWeapon:  model.MasteryLevel(record[11]),
			MartialWeapon: model.MasteryLevel(record[12]),
		}

		if record[13] != "" {
			tradition, _ := a.DB.GetTraditionByName(record[13])
			characterClass.TraditionID = &tradition.ID
		}

		characterClasses = append(characterClasses, characterClass)
		if existCharacterClass, err := a.DB.GetCharacterClassByName(characterClass.Name); err == nil && existCharacterClass != nil {
			continue
		}
		err = a.DB.CreateCharacterClass(&characterClass)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *LoadCSVApi) LoadFeat(ctx *gin.Context) {
	file, err := os.Open("./csv/Feat.csv")
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
	var feats []model.Feat

	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		log.Println(record)

		if len(record) != 8 {
			log.Printf("Wrong record count %v", record)
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		level, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}

		skill, err := a.DB.GetSkillByName(record[5])
		if err != nil {
			log.Fatal(err)
		}

		traits := a.GetTraits(ctx, record[6])

		feat := model.Feat{
			Name:             record[0],
			Description:      record[1],
			Level:            uint8(level),
			Rarity:           model.Rarity(record[3]),
			PrerequisiteFeat: &record[7],
		}

		if skill != nil {
			feat.PrerequisiteSkillID = &skill.ID
			feat.PrerequisiteMastery = model.MasteryLevel(record[4])
		}

		if traits != nil {
			featTraits, err := a.DB.FindTraits(traits)
			if err != nil {
				log.Fatal(err)
			}
			feat.Traits = featTraits
		}

		feats = append(feats, feat)
		if existFeat, err := a.DB.GetFeatByName(feat.Name); err == nil && existFeat != nil {
			continue
		}
		err = a.DB.CreateFeat(&feat)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *LoadCSVApi) GetTraits(ctx *gin.Context, traits string) []uint {
	parts := strings.Split(traits, ", ")
	var traitsID []uint
	for _, part := range parts {
		log.Println(part)
		trait, err := a.DB.GetTraitByName(part)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil
		}
		log.Println(trait)
		log.Println(trait.ID)
		traitsID = append(traitsID, trait.ID)
	}
	if len(traitsID) == 0 {
		return nil
	}
	return traitsID
}

func (a *LoadCSVApi) GetTraditions(ctx *gin.Context, traditions string) []uint {
	parts := strings.Split(traditions, ", ")
	var traditionsID []uint
	for _, part := range parts {
		tradition, err := a.DB.GetTraditionByName(part)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil
		}
		traditionsID = append(traditionsID, tradition.ID)
	}
	if len(traditionsID) == 0 {
		return nil
	}
	return traditionsID
}
