package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kingdom/api"
	"kingdom/auth"
	"kingdom/config"
	"kingdom/consumer"
	"kingdom/database"
	"kingdom/docs"
	gerror "kingdom/error"
)

func Create(db *database.GormDatabase, conf *config.Configuration, consumer *consumer.RMQConsumer) (*gin.Engine, func()) {
	g := gin.New()
	g.RemoteIPHeaders = []string{"X-Forwarded-For"}
	g.SetTrustedProxies(conf.Server.TrustedProxies)
	g.ForwardedByClientIP = true
	g.Use(func(ctx *gin.Context) {
		if ctx.Request.RemoteAddr == "@" {
			ctx.Request.RemoteAddr = "localhost:8080"
		}
	})
	authentication := auth.Auth{DB: db}

	userChangeNotifier := new(api.UserChangeNotifier)
	userHandler := api.UserApi{
		DB:                 db,
		PasswordStrength:   conf.PassStrength,
		UserChangeNotifier: userChangeNotifier,
		Registration:       conf.Registration}

	characterHandler := api.CharacterApi{DB: db}
	characterClassHandler := api.CharacterClassApi{DB: db}
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

	authHandler := api.Controller{DB: db}

	userRabbitHandler := api.UserRabbitApi{DB: db, Consumer: consumer}

	g.NoRoute(gerror.NotFound())

	g.Use(cors.New(auth.CorsConfig(conf)))

	docs.SwaggerInfo.BasePath = ""
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.POST("/user", userHandler.CreateUser)
	g.POST("/login", authHandler.Login)
	g.GET("/validate", authHandler.Validate)

	g.POST("/a/rabbit", userRabbitHandler.CreateUserRabbit)

	userGroup := g.Group("/user").Use(authentication.RequireJWT)
	{
		userGroup.GET("", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.DELETE("/:id", userHandler.DeleteUserByID)
		userGroup.PATCH("/:id", userHandler.UpdateUser)
		userGroup.PATCH("/password", userHandler.ChangePassword)
	}
	characterGroup := g.Group("/character").Use(authentication.RequireJWT)
	{
		characterGroup.POST("", characterHandler.CreateCharacter)
		characterGroup.GET("/:id", characterHandler.GetCharacterByID)
		characterGroup.GET("", characterHandler.GetCharacters)
		characterGroup.PATCH("/:id", characterHandler.UpdateCharacter)
		characterGroup.DELETE("/:id", characterHandler.DeleteCharacter)
	}
	g.POST("/character_feat", characterHandler.AddCharacterFeat)
	godGroup := g.Group("/god").Use(authentication.RequireAdmin)
	{
		godGroup.POST("", godHandler.CreateGod)
		godGroup.PATCH("/:id", godHandler.UpdateGod)
		godGroup.DELETE("/:id", godHandler.DeleteGod)
	}
	godGroup.GET("/:id", godHandler.GetGodById).Use(authentication.RequireJWT)
	godGroup.GET("", godHandler.GetGods).Use(authentication.RequireJWT)

	domainGroup := g.Group("/domain").Use(authentication.RequireAdmin)
	{
		domainGroup.POST("", domainHandler.CreateDomain)
		domainGroup.PATCH("/:id", domainHandler.UpdateDomain)
		domainGroup.DELETE("/:id", domainHandler.DeleteDomain)
	}
	g.GET("/domain/:id", domainHandler.GetDomainByID).Use(authentication.RequireJWT)
	g.GET("/domain", domainHandler.GetDomains).Use(authentication.RequireJWT)

	featGroup := g.Group("/feat").Use(authentication.RequireJWT)
	{
		featGroup.POST("", featHandler.CreateFeat)
		featGroup.PATCH("/:id", featHandler.UpdateFeat)
		featGroup.GET("", featHandler.GetFeats)
		featGroup.GET("/:id", featHandler.GetFeatByID)
		featGroup.DELETE("/:id", featHandler.DeleteFeat)
	}

	raceGroup := g.Group("/race").Use(authentication.RequireJWT)
	{
		raceGroup.POST("", raceHandler.CreateRace)
		raceGroup.PATCH("/:id", raceHandler.UpdateRace)
		raceGroup.GET("", raceHandler.GetRaces)
		raceGroup.GET("/:id", raceHandler.GetRaceByID)
		raceGroup.DELETE("/:id", raceHandler.DeleteRace)
	}

	ancestryGroup := g.Group("/ancestry").Use(authentication.RequireJWT)
	{
		ancestryGroup.POST("", ancestryHandler.CreateAncestry)
		ancestryGroup.PATCH("/:id", ancestryHandler.UpdateAncestry)
		ancestryGroup.GET("", ancestryHandler.GetAncestries)
		ancestryGroup.GET("/:id", ancestryHandler.GetAncestryByID)
		ancestryGroup.DELETE("/:id", ancestryHandler.DeleteAncestry)
	}

	actionGroup := g.Group("/action").Use(authentication.RequireJWT)
	{
		actionGroup.POST("", actionHandler.CreateAction)
		actionGroup.PATCH("/:id", actionHandler.UpdateAction)
		actionGroup.GET("", actionHandler.GetActions)
		actionGroup.GET("/:id", actionHandler.GetActionByID)
		actionGroup.DELETE("/:id", actionHandler.DeleteAction)
	}

	backgroundGroup := g.Group("/background").Use(authentication.RequireJWT)
	{
		backgroundGroup.POST("", backgroundHandler.CreateBackground)
		backgroundGroup.PATCH("/:id", backgroundHandler.UpdateBackground)
		backgroundGroup.GET("/:id", backgroundHandler.GetBackgroundByID)
		backgroundGroup.GET("", backgroundHandler.GetBackgrounds)
		backgroundGroup.DELETE("/:id", backgroundHandler.DeleteBackground)
	}

	traditionGroup := g.Group("/tradition").Use(authentication.RequireJWT)
	{
		traditionGroup.POST("", traditionHandler.CreateTradition)
		traditionGroup.PATCH("/:id", traditionHandler.UpdateTradition)
		traditionGroup.GET("", traditionHandler.GetTraditions)
		traditionGroup.GET("/:id", traditionHandler.GetTraditionByID)
		traditionGroup.DELETE("/:id", traditionHandler.DeleteTradition)
	}

	traitGroup := g.Group("/trait").Use(authentication.RequireJWT)
	{
		traitGroup.POST("", traitHandler.CreateTrait)
		traitGroup.PATCH("/:id", traitHandler.UpdateTrait)
		traitGroup.GET("", traitHandler.GetTraits)
		traitGroup.GET("/:id", traitHandler.GetTraitByID)
		traitGroup.DELETE("/:id", traitHandler.DeleteTrait)
	}

	characterClassGroup := g.Group("/class").Use(authentication.RequireJWT)
	{
		characterClassGroup.POST("", characterClassHandler.CreateCharacterClass)
	}
	itemGroup := g.Group("/item").Use(authentication.RequireJWT)
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

	characterItemGroup := g.Group("/character-item").Use(authentication.RequireJWT)
	{
		characterItemGroup.POST("", characterItemHandler.CreateCharacterItem)
		characterItemGroup.GET("/:id", characterItemHandler.GetCharacterItemByID)
		characterItemGroup.GET("/list/:character_id", characterItemHandler.GetCharacterItems)
		characterItemGroup.DELETE("/:id", characterItemHandler.DeleteCharacterItem)
		characterItemGroup.PATCH("/:id", characterItemHandler.UpdateCharacterItem)
	}

	characterSkillGroup := g.Group("/character-skill").Use(authentication.RequireJWT)
	{
		characterSkillGroup.POST("", characterSkillHandler.CharacterSkillCreate)
		characterSkillGroup.GET("", characterSkillHandler.GetCharacterSkills)
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
