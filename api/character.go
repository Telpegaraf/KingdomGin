package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/auth"
	"kingdom/model"
	"net/http"
)

type CharacterDatabase interface {
	GetCharacterByID(id uint) (*model.Character, error)
	CreateCharacter(character *model.Character) error
	GetCharacters(id uint) ([]*model.Character, error)
	UpdateCharacter(character *model.Character) error
	DeleteCharacterByID(id uint) error
	GetUserByID(id uint) (*model.User, error)
	CreateAttribute(stat *model.Attribute) error
	CreateSlot(slot *model.Slot) error
	CreateCharacterBoost(stat *model.CharacterBoost) error
	CreateCharacterDefence(characterDefence *model.CharacterDefence) error
	GetRaceByID(id uint) (*model.Race, error)
	GetCharacterClassByID(id uint) (*model.CharacterClass, error)
	GetCharacterDefenceByID(id uint) (*model.CharacterDefence, error)
	UpdateHitPoint(defence *model.CharacterDefence) error
	GetSkills() ([]*model.Skill, error)
	CharacterSkillCreate(characterSkill *model.CharacterSkill) error
	GetBackgroundByID(id uint) (*model.Background, error)
	CreateCharacterFeat(characterFeat *model.CharacterFeat) error
	GetFeatByID(id uint) (*model.Feat, error)
	GetCharacterSkillByCharacterID(characterId uint, skillID uint) (*model.CharacterSkill, error)
	CreateCharacterInfo(*model.CharacterInfo) error
	GetAncestryByID(id uint) (*model.Ancestry, error)
}

type CharacterApi struct {
	DB CharacterDatabase
}

// GetCharacterByID godoc
//
// @Summary Returns Character by id
// @Description Retrieve Character details using its ID
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "character id"
// @Success 200 {object} model.CharacterExternal "character details"
// @Failure 404 {string} string "Character not found"
// @Router /character/{id} [get]
func (a *CharacterApi) GetCharacterByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		character, err := a.DB.GetCharacterByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if character != nil {
			ctx.JSON(http.StatusOK, ToExternalCharacter(character))
		} else {
			ctx.JSON(404, gin.H{"error": "Character not found"})
		}
	})
}

// GetCharacters godoc
//
// @Summary Returns all characters
// @Description Return all characters for current user
// @Tags Character
// @Accept json
// @Produce json
// @Success 200 {object} model.CharacterExternal "Character details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /character [get]
func (a *CharacterApi) GetCharacters(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	characters, err := a.DB.GetCharacters(userID.(uint))

	if success := SuccessOrAbort(ctx, 500, err); !success {
		return
	}
	var resp []*model.CharacterExternal
	for _, character := range characters {
		resp = append(resp, ToExternalCharacter(character))
	}
	ctx.JSON(http.StatusOK, resp)
}

// CreateCharacter godoc
//
// @Summary Create and returns character or nil
// @Description Create new character
// @Tags Character
// @Accept json
// @Produce json
// @Param character body model.CreateCharacter true "Character data"
// @Success 201 {object} model.CharacterExternal "Character details"
// @Failure 401 {string} string "Unauthorized"
// @Router /character [post]
func (a *CharacterApi) CreateCharacter(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	character := &model.CreateCharacter{}
	if err := ctx.Bind(character); err == nil {
		internal := &model.Character{
			Name:             character.Name,
			Alias:            character.Alias,
			LastName:         character.LastName,
			UserID:           userID.(uint),
			CharacterClassID: character.CharacterClassID,
			AncestryID:       character.AncestryID,
			BackgroundID:     character.BackgroundID,
			RaceID:           character.RaceID,
		}

		ancestry, err := a.DB.GetAncestryByID(character.AncestryID)
		if err != nil || ancestry.RaceID != character.RaceID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Ancestry or Race ID"})
			return
		}

		if success := SuccessOrAbort(ctx, 500, a.DB.CreateCharacter(internal)); !success {
			return
		}
		newCharacter, _ := a.DB.GetCharacterByID(internal.ID)
		go func() {
			race, _ := a.DB.GetRaceByID(character.RaceID)
			characterClass, _ := a.DB.GetCharacterClassByID(character.CharacterClassID)
			background, _ := a.DB.GetBackgroundByID(character.BackgroundID)
			a.CreateAttribute(ctx, newCharacter.ID, race)
			a.CreateSlot(ctx, internal.ID)
			a.CreateCharacterBoost(ctx, newCharacter.ID, race)
			a.CreateSkills(ctx, newCharacter)
			a.CreateCharacterDefence(ctx, newCharacter.ID, race, characterClass)
			a.CreateCharacterFeat(newCharacter.ID, background)
		}()
		ctx.JSON(http.StatusCreated, ToExternalCharacter(newCharacter))
	}
}

// UpdateCharacter Updates Character by ID
//
// @Summary Updates Character by ID or nil
// @Description Permissions for Character's User or Admin
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "Character id"
// @Param character body model.CharacterUpdate true "Character data"
// @Success 200 {object} model.CharacterExternal "Character details"
// @Failure 404 {string} string "Character doesn't exist"
// @Router /character/{id} [patch]
func (a *CharacterApi) UpdateCharacter(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))

	withID(ctx, "id", func(id uint) {
		var character *model.CharacterUpdate
		if err := ctx.Bind(&character); err == nil {
			oldCharacter, err := a.DB.GetCharacterByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldCharacter != nil {
				if oldCharacter.UserID != user.ID && !user.Admin {
					ctx.JSON(http.StatusForbidden, gin.H{"error": "You can't access for this API"})
					return
				}
				internal := &model.Character{
					ID:               oldCharacter.ID,
					CharacterClassID: oldCharacter.CharacterClassID,
					Name:             character.Name,
					Alias:            character.Alias,
					LastName:         character.LastName,
					UserID:           oldCharacter.UserID,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateCharacter(internal)); success {
					return
				}
				ctx.JSON(http.StatusOK, ToExternalCharacter(internal))
				if character.Level != oldCharacter.Level {
					go func() {
						a.ChangeHitPoint(ctx, internal, character.Level-oldCharacter.Level)
					}()
				}
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func (a *CharacterApi) ChangeHitPoint(ctx *gin.Context, character *model.Character, levelCount int8) {
	characterClass, _ := a.DB.GetCharacterClassByID(character.CharacterClassID)
	characterDefence, _ := a.DB.GetCharacterDefenceByID(character.ID)
	maxHtPoint := characterDefence.MaxHitPoint + (characterClass.HitPoint)*uint16(levelCount)
	hitPoint := characterDefence.HitPoint
	if maxHtPoint < characterDefence.HitPoint {
		hitPoint = maxHtPoint
	}
	internal := &model.CharacterDefence{
		ID:          character.ID,
		MaxHitPoint: maxHtPoint,
		HitPoint:    hitPoint,
	}
	err := a.DB.UpdateHitPoint(internal)
	if err != nil {
		return
	}

}

// DeleteCharacter Deletes Character by ID
//
// @Summary Deletes Character by ID or returns nil
// @Description Permissions for Character's User or Admin
// @Tags Character
// @Accept json
// @Produce json
// @Param id path int true "Character id"
// @Success 204
// @Failure 404 {string} string "Character doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /character/{id} [delete]
func (a *CharacterApi) DeleteCharacter(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))

	withID(ctx, "id", func(id uint) {
		character, err := a.DB.GetCharacterByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if character != nil {
			if character.UserID != user.ID && !user.Admin {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "You can't access for this API"})
				return
			}
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteCharacterByID(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}

func ToExternalCharacter(character *model.Character) *model.CharacterExternal {
	return &model.CharacterExternal{
		ID:               character.ID,
		Name:             character.Name,
		Alias:            character.Alias,
		LastName:         character.LastName,
		UserID:           character.UserID,
		Level:            character.Level,
		CharacterItem:    character.CharacterItem,
		CharacterBoost:   character.Boost,
		Attribute:        character.Attribute,
		Slot:             character.Slot,
		CharacterClassID: character.CharacterClassID,
		RaceID:           character.RaceID,
		AncestryID:       character.AncestryID,
		BackgroundID:     character.BackgroundID,
	}
}

func (a *CharacterApi) CreateSkills(ctx *gin.Context, character *model.Character) {
	backgroundCharacter, _ := a.DB.GetBackgroundByID(character.BackgroundID)
	skills, _ := a.DB.GetSkills()
	for _, skill := range skills {
		mastery := model.None
		if *backgroundCharacter.FirstSkillID == skill.ID || *backgroundCharacter.SecondSkillID == skill.ID {
			mastery = model.Train
		}
		characterSkill := &model.CharacterSkill{
			CharacterID: character.ID,
			SkillID:     skill.ID,
			Mastery:     mastery,
		}
		err := a.DB.CharacterSkillCreate(characterSkill)
		if err != nil {
			return
		}
	}
}
