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
)

type CharacterClassDatabase interface {
	CreateCharacterClass(characterClass *model.CharacterClass) error
	GetCharacterClassByName(name string) (*model.CharacterClass, error)
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

// LoadCharacterClass godoc
//
// @Summary Create Character Class from csv file on server or nil
// @Description Permissions for Admin
// @Tags Character Class
// @Accept json
// @Produce json
// @Success 201 {object} model.CharacterClassExternal "Tradition details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /class/load [post]
func (a *CharacterClassApi) LoadCharacterClass(ctx *gin.Context) {
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
