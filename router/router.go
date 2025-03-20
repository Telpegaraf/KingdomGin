package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kingdom/api"
	"kingdom/auth"
	"kingdom/config"
	"kingdom/database"
	"kingdom/docs"
	gerror "kingdom/error"
	"time"
)

func Create(db *database.GormDatabase, conf *config.Configuration) (*gin.Engine, func()) {
	g := gin.New()

	g.RemoteIPHeaders = []string{"X-Forwarded-For"}
	err := g.SetTrustedProxies(conf.Server.TrustedProxies)
	if err != nil {
		return nil, nil
	}
	g.ForwardedByClientIP = true
	g.Use(func(ctx *gin.Context) {
		if ctx.Request.RemoteAddr == "@" {
			ctx.Request.RemoteAddr = "localhost:8080"
		}
	})

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://kingdom-p2e.ru"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "x-initData"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authentication := auth.Auth{DB: db}

	userHandler := api.UserApi{
		DB:               db,
		PasswordStrength: conf.PassStrength,
		Registration:     conf.Registration,
	}

	characterHandler := api.CharacterApi{DB: db}
	characterClassHandler := api.CharacterClassApi{DB: db}
	classFeatureHandler := api.ClassFeatureApi{DB: db}
	itemHandler := api.ItemApi{DB: db}
	characterItemHandler := api.CharacterItemApi{DB: db}
	slotHandler := api.SlotApi{DB: db}
	attributeHandler := api.AttributeApi{DB: db}
	characterBoostHandler := api.CharacterBoostApi{DB: db}
	godHandler := api.GodApi{DB: db}
	domainHandler := api.DomainApi{DB: db}
	featHandler := api.FeatAPI{DB: db}
	raceHandler := api.RaceApi{DB: db}
	ancestryHandler := api.AncestryApi{DB: db}
	traditionHandler := api.TraditionApi{DB: db}
	actionHandler := api.ActionApi{DB: db}
	traitHandler := api.TraitApi{DB: db}
	characterSkillHandler := api.CharacterSkillApi{DB: db}
	backgroundHandler := api.BackgroundApi{DB: db}
	skillHandler := api.SkillApi{DB: db}
	spellHandler := api.SpellAPI{DB: db}
	loadCSVHandler := api.LoadCSVApi{DB: db}

	g.NoRoute(gerror.NotFound())

	g.Use(cors.New(auth.CorsConfig(conf)))

	docs.SwaggerInfo.BasePath = ""
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.POST("/user", userHandler.CreateUser)

	apiGroup := g.Group("/api")
	//apiGroup.Use(middleware.CheckWebAppSignatureMiddleware())

	adminGroup := apiGroup.Group("/admin")
	{
		adminGroup.POST("/csv", loadCSVHandler.LoadCSV)
	}

	userGroup := apiGroup.Group("/user").Use(authentication.RequireJWT)
	{
		userGroup.GET("", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.DELETE("/:id", userHandler.DeleteUserByID)
	}
	characterGroup := apiGroup.Group("/character")
	{
		characterGroup.POST("", characterHandler.CreateCharacter)
		characterGroup.GET("/:id", characterHandler.GetCharacterByID)
		characterGroup.GET("", characterHandler.GetCharacters)
		characterGroup.PATCH("/:id", characterHandler.UpdateCharacter)
		characterGroup.DELETE("/:id", characterHandler.DeleteCharacter)
	}
	g.POST("/character_feat", characterHandler.AddCharacterFeat)
	godGroup := apiGroup.Group("/god").Use(authentication.RequireAdmin)
	{
		godGroup.POST("", godHandler.CreateGod)
		godGroup.PATCH("/:id", godHandler.UpdateGod)
		godGroup.DELETE("/:id", godHandler.DeleteGod)
	}
	godGroup.GET("/:id", godHandler.GetGodById).Use(authentication.RequireJWT)
	godGroup.GET("", godHandler.GetGods).Use(authentication.RequireJWT)

	domainGroup := apiGroup.Group("/domain").Use(authentication.RequireAdmin)
	{
		domainGroup.POST("", domainHandler.CreateDomain)
		domainGroup.POST("/load", domainHandler.LoadDomain)
		domainGroup.PATCH("/:id", domainHandler.UpdateDomain)
		domainGroup.DELETE("/:id", domainHandler.DeleteDomain)
	}
	g.GET("/domain/:id", domainHandler.GetDomainByID).Use(authentication.RequireJWT)
	g.GET("/domain", domainHandler.GetDomains).Use(authentication.RequireJWT)

	skillGroup := apiGroup.Group("/skill").Use(authentication.RequireAdmin)
	{
		skillGroup.POST("", skillHandler.CreateSkill)
		skillGroup.PATCH("/:id", skillHandler.UpdateSkill)
		skillGroup.DELETE("/:id", skillHandler.DeleteSkill)
	}
	g.GET("/skill", skillHandler.GetSkills).Use(authentication.RequireJWT)
	g.GET("/skill/:id", skillHandler.GetSkillByID).Use(authentication.RequireJWT)

	featGroup := apiGroup.Group("/feat").Use(authentication.RequireAdmin)
	{
		featGroup.POST("", featHandler.CreateFeat)
		featGroup.PATCH("/:id", featHandler.UpdateFeat)
		featGroup.DELETE("/:id", featHandler.DeleteFeat)
	}
	g.GET("/feat", featHandler.GetFeats).Use(authentication.RequireJWT)
	g.GET("/feat/:id", featHandler.GetFeatByID).Use(authentication.RequireJWT)

	raceGroup := apiGroup.Group("/race").Use(authentication.RequireAdmin)
	{
		raceGroup.POST("", raceHandler.CreateRace)
		raceGroup.PATCH("/:id", raceHandler.UpdateRace)
		raceGroup.DELETE("/:id", raceHandler.DeleteRace)
	}
	g.GET("/race", raceHandler.GetRaces).Use(authentication.RequireJWT)
	g.GET("/race/:id", raceHandler.GetRaceByID).Use(authentication.RequireJWT)

	spellGroup := apiGroup.Group("/spell").Use(authentication.RequireAdmin)
	{
		spellGroup.POST("", spellHandler.CreateSpell)
		spellGroup.PATCH("/:id", spellHandler.UpdateSpell)
		spellGroup.DELETE("/:id", spellHandler.DeleteSpell)
	}
	g.GET("/spell", spellHandler.GetSpells).Use(authentication.RequireJWT)
	g.GET("/spell/:id", spellHandler.GetSpellByID).Use(authentication.RequireJWT)

	ancestryGroup := apiGroup.Group("/ancestry").Use(authentication.RequireAdmin)
	{
		ancestryGroup.POST("", ancestryHandler.CreateAncestry)
		ancestryGroup.PATCH("/:id", ancestryHandler.UpdateAncestry)
		ancestryGroup.DELETE("/:id", ancestryHandler.DeleteAncestry)
	}
	g.GET("/ancestry", ancestryHandler.GetAncestries).Use(authentication.RequireJWT)
	g.GET("/ancestry/:id", ancestryHandler.GetAncestryByID).Use(authentication.RequireJWT)

	actionGroup := apiGroup.Group("/action").Use(authentication.RequireAdmin)
	{
		actionGroup.POST("", actionHandler.CreateAction)
		actionGroup.PATCH("/:id", actionHandler.UpdateAction)
		actionGroup.DELETE("/:id", actionHandler.DeleteAction)
	}
	g.GET("/action", actionHandler.GetActions).Use(authentication.RequireJWT)
	g.GET("/action/:id", actionHandler.GetActionByID).Use(authentication.RequireJWT)

	backgroundGroup := apiGroup.Group("/background").Use(authentication.RequireAdmin)
	{
		backgroundGroup.POST("", backgroundHandler.CreateBackground)
		backgroundGroup.PATCH("/:id", backgroundHandler.UpdateBackground)
		backgroundGroup.DELETE("/:id", backgroundHandler.DeleteBackground)
	}
	g.GET("/background", backgroundHandler.GetBackgrounds).Use(authentication.RequireJWT)
	g.GET("/background/:id", backgroundHandler.GetBackgroundByID).Use(authentication.RequireJWT)

	traditionGroup := apiGroup.Group("/tradition").Use(authentication.RequireAdmin)
	{
		traditionGroup.POST("", traditionHandler.CreateTradition)
		traditionGroup.PATCH("/:id", traditionHandler.UpdateTradition)
		traditionGroup.DELETE("/:id", traditionHandler.DeleteTradition)
	}
	g.GET("/tradition", traditionHandler.GetTraditions).Use(authentication.RequireJWT)
	g.GET("/tradition/:id", traditionHandler.GetTraditionByID).Use(authentication.RequireJWT)

	traitGroup := apiGroup.Group("/trait").Use(authentication.RequireAdmin)
	{
		traitGroup.POST("", traitHandler.CreateTrait)
		traitGroup.PATCH("/:id", traitHandler.UpdateTrait)
		traitGroup.DELETE("/:id", traitHandler.DeleteTrait)
	}
	g.GET("/trait", traitHandler.GetTraits).Use(authentication.RequireJWT)
	g.GET("/trait/:id", traitHandler.GetTraitByID).Use(authentication.RequireJWT)

	characterClassGroup := apiGroup.Group("/class").Use(authentication.RequireAdmin)
	{
		characterClassGroup.POST("", characterClassHandler.CreateCharacterClass)
		characterClassGroup.PATCH("/:id", characterClassHandler.UpdateCharacterClass)
		characterClassGroup.DELETE("/:id", characterClassHandler.DeleteCharacterClass)
	}
	g.GET("/class", characterClassHandler.GetCharacterClasses).Use(authentication.RequireJWT)
	g.GET("/class/:id", characterClassHandler.GetCharacterClassByID).Use(authentication.RequireJWT)

	g.GET("/class-feature/:id", classFeatureHandler.GetClassFeatureByID).Use(authentication.RequireAdmin)
	g.GET("/class-feature/all/:id", classFeatureHandler.GetAllFeature).Use(authentication.RequireAdmin)
	g.GET("/skill-feature/:id", classFeatureHandler.GetClassSkillFeatureByID).Use(authentication.RequireJWT)
	itemGroup := apiGroup.Group("/item").Use(authentication.RequireJWT)
	{
		itemGroup.GET("", itemHandler.GetItems)
		itemGroup.GET(":id", itemHandler.GetItemByID)
		itemGroup.GET("/armor", itemHandler.GetArmors)
		itemGroup.GET("/armor/:id", itemHandler.GetArmorByID)
		itemGroup.GET("/weapon", itemHandler.GetWeapons)
		itemGroup.GET("/weapon/:id", itemHandler.GetWeaponByID)
		itemGroup.GET("/gear", itemHandler.GetGears)
		itemGroup.GET("/gear/:id", itemHandler.GetGearByID)
	}
	itemGroup.DELETE("/:id", itemHandler.DeleteItem).Use(authentication.RequireAdmin)
	itemGroup.POST("/armor", itemHandler.CreateArmor).Use(authentication.RequireAdmin)
	itemGroup.PATCH("/armor/:id", itemHandler.UpdateArmor).Use(authentication.RequireAdmin)
	itemGroup.POST("/weapon", itemHandler.CreateWeapon).Use(authentication.RequireAdmin)
	itemGroup.PATCH("/weapon/:id", itemHandler.UpdateWeapon).Use(authentication.RequireAdmin)
	itemGroup.POST("/gear", itemHandler.CreateGear).Use(authentication.RequireAdmin)
	itemGroup.PATCH("/gear/:id", itemHandler.UpdateGear).Use(authentication.RequireAdmin)

	characterItemGroup := apiGroup.Group("/character-item").Use(authentication.RequireJWT)
	{
		characterItemGroup.POST("", characterItemHandler.CreateCharacterItem)
		characterItemGroup.GET("/:id", characterItemHandler.GetCharacterItemByID)
		characterItemGroup.GET("/list/:character_id", characterItemHandler.GetCharacterItems)
		characterItemGroup.DELETE("/:id", characterItemHandler.DeleteCharacterItem)
		characterItemGroup.PATCH("/:id", characterItemHandler.UpdateCharacterItem)
	}

	characterSkillGroup := apiGroup.Group("/character-skill").Use(authentication.RequireJWT)
	{
		characterSkillGroup.POST("", characterSkillHandler.CharacterSkillCreate)
		characterSkillGroup.GET("/:id", characterSkillHandler.GetCharacterSkills)
		characterSkillGroup.PATCH("/:id", characterSkillHandler.UpdateCharacterSkill)
	}

	g.GET("/slot/:id", slotHandler.GetSlotByID).Use(authentication.RequireJWT)
	g.PATCH("/slot/:id", slotHandler.UpdateSlot).Use(authentication.RequireJWT)

	g.GET("/attribute/:id", attributeHandler.GetAttributeByID).Use(authentication.RequireJWT)
	g.PATCH("/attribute/:id", attributeHandler.UpdateAttribute).Use(authentication.RequireJWT)

	g.GET("/character_boost/:id", characterBoostHandler.GetCharacterBoostByID).Use(authentication.RequireJWT)
	g.PATCH("/character_boost/:id", characterBoostHandler.UpdateCharacterBoost).Use(authentication.RequireJWT)

	return g, func() {}
}
