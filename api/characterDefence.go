package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type CharacterDefenceDatabase interface {
	GetCharacterDefenceByID(id uint) (*model.CharacterDefence, error)
	UpdateCharacterDefence(CharacterBoosts *model.CharacterDefence) error
}

type CharacterDefenceApi struct {
	DB CharacterDefenceDatabase
}

func (a *CharacterApi) CreateCharacterDefence(ctx *gin.Context,
	characterId uint,
	race *model.Race,
	characterClass *model.CharacterClass) {
	internal := &model.CharacterDefence{
		CharacterID: characterId,
		HitPoint:    race.HitPoint + characterClass.HitPoint,
		MaxHitPoint: race.HitPoint + characterClass.HitPoint,
		Perception:  characterClass.Perception,
		Fortitude:   characterClass.Fortitude,
		Reflex:      characterClass.Reflex,
		Will:        characterClass.Will,
		Unarmed:     characterClass.UnarmedArmor,
		LightArmor:  characterClass.LightArmor,
		MediumArmor: characterClass.MediumArmor,
		HeavyArmor:  characterClass.HeavyArmor,
	}
	a.DB.CreateCharacterDefence(internal)
}

func (a *CharacterDefenceApi) GetCharacterDefence(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		characterDefence, err := a.DB.GetCharacterDefenceByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if characterDefence != nil {
			ctx.JSON(http.StatusOK, ToExternalCharacterDefence(characterDefence))
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "character defence not found"})
		}
	})
}

func ToExternalCharacterDefence(characterDefence *model.CharacterDefence) *model.CharacterDefenceExternal {
	return &model.CharacterDefenceExternal{
		CharacterID:       characterDefence.CharacterID,
		ArmorClass:        characterDefence.ArmorClass,
		MaxHitPoint:       characterDefence.MaxHitPoint,
		TemporaryHitPoint: characterDefence.TemporaryHitPoint,
		Dying:             characterDefence.Dying,
		Wounded:           characterDefence.Wounded,
		Speed:             characterDefence.Speed,
		HitPoint:          characterDefence.HitPoint,
		Perception:        characterDefence.Perception,
		Fortitude:         characterDefence.Fortitude,
		Reflex:            characterDefence.Reflex,
		Will:              characterDefence.Will,
		Unarmed:           characterDefence.Unarmed,
		LightArmor:        characterDefence.LightArmor,
		MediumArmor:       characterDefence.MediumArmor,
		HeavyArmor:        characterDefence.HeavyArmor,
	}
}
